// Copyright 2017 Istio Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package kubernetes provides functionality to adapt mixer behavior to the
// kubernetes environment. Primarily, it is used to generate values as part
// of Mixer's attribute generation preprocessing phase. These values will be
// transformed into attributes that can be used for subsequent config
// resolution and adapter dispatch and execution.
package kubernetes

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"k8s.io/api/core/v1"
	k8s "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // needed for auth
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	"context"

	"istio.io/istio/mixer/adapter/kubernetes2/config"
	adapter_template_kubernetes "istio.io/istio/mixer/adapter/kubernetes2/template"
	"istio.io/istio/mixer/pkg/adapter"
)

type (
	builder struct {
		adapterConfig        *config.Params
		newCacheControllerFn controllerFactoryFn
		needsCacheInit       bool
	}

	handler struct {
		pods   cacheController
		log    adapter.Logger
		params config.Params
	}

	// used strictly for testing purposes
	controllerFactoryFn func(kubeconfigPath string, refreshDuration time.Duration, env adapter.Env) (cacheController, error)
)

var _ adapter_template_kubernetes.Handler = &handler{}

func (h *handler) Close() error {
	return nil
}

func (h *handler) GenerateKubernetesAttributes(context.Context, *adapter_template_kubernetes.Instance) (*adapter_template_kubernetes.Output, error) {
	return nil, nil
}

var _ adapter_template_kubernetes.HandlerBuilder = &builder{}

// SetAdapterConfig gives the builder the adapter-level configuration state.
func (b *builder) SetAdapterConfig(c adapter.Config) {
	b.adapterConfig = c.(*config.Params)
}

// Validate is responsible for ensuring that all the configuration state given to the builder is
// correct. The Build method is only invoked when Validate has returned success.
func (b *builder) Validate() (ce *adapter.ConfigErrors) {
	params := b.adapterConfig
	if len(params.SourceUidInputName) == 0 {
		ce = ce.Appendf("sourceUidInputName", "field must be populated")
	}
	if len(params.DestinationUidInputName) == 0 {
		ce = ce.Appendf("destinationUidInputName", "field must be populated")
	}
	if len(params.OriginUidInputName) == 0 {
		ce = ce.Appendf("originUidInputName", "field must be populated")
	}
	if len(params.SourceIpInputName) == 0 {
		ce = ce.Appendf("sourceIpInputName", "field must be populated")
	}
	if len(params.DestinationIpInputName) == 0 {
		ce = ce.Appendf("destinationIpInputName", "field must be populated")
	}
	if len(params.OriginIpInputName) == 0 {
		ce = ce.Appendf("originIpInputName", "field must be populated")
	}
	if len(params.SourcePrefix) == 0 {
		ce = ce.Appendf("sourcePrefix", "field must be populated")
	}
	if len(params.DestinationPrefix) == 0 {
		ce = ce.Appendf("destinationPrefix", "field must be populated")
	}
	if len(params.OriginPrefix) == 0 {
		ce = ce.Appendf("originPrefix", "field must be populated")
	}
	if len(params.LabelsValueName) == 0 {
		ce = ce.Appendf("labelsValueName", "field must be populated")
	}
	if len(params.PodIpValueName) == 0 {
		ce = ce.Appendf("podIpValueName", "field must be populated")
	}
	if len(params.PodNameValueName) == 0 {
		ce = ce.Appendf("podNameValueName", "field must be populated")
	}
	if len(params.HostIpValueName) == 0 {
		ce = ce.Appendf("hostIpValueName", "field must be populated")
	}
	if len(params.NamespaceValueName) == 0 {
		ce = ce.Appendf("namespaceValueName", "field must be populated")
	}
	if len(params.ServiceAccountValueName) == 0 {
		ce = ce.Appendf("serviceAccountValueName", "field must be populated")
	}
	if len(params.ServiceValueName) == 0 {
		ce = ce.Appendf("serviceValueName", "field must be populated")
	}
	if len(params.PodLabelForService) == 0 {
		ce = ce.Appendf("podLabelForService", "field must be populated")
	}
	if len(params.PodLabelForIstioComponentService) == 0 {
		ce = ce.Appendf("podLabelForIstioComponentService", "field must be populated")
	}
	if len(params.FullyQualifiedIstioIngressServiceName) == 0 {
		ce = ce.Appendf("fullyQualifiedIstioIngressServiceName", "field must be populated")
	}
	if len(params.ClusterDomainName) == 0 {
		ce = ce.Appendf("clusterDomainName", "field must be populated")
	} else if len(strings.Split(params.ClusterDomainName, ".")) != 3 {
		ce = ce.Appendf("clusterDomainName", "must have three segments, separated by '.' ('svc.cluster.local', for example)")
	}
	return
}

// Build must return a handler that implements all the template-specific runtime request serving
// interfaces that the Builder was configured for.
// This means the Handler returned by the Build method must implement all the runtime interfaces for all the
// template the Adapter supports.
// If the returned Handler fails to implement the required interface that builder was registered for, Mixer will
// report an error and stop serving runtime traffic to the particular Handler.
func (b *builder) Build(ctx context.Context, env adapter.Env) (adapter.Handler, error) {
	paramsProto := b.adapterConfig
	stopChan := make(chan struct{})
	var pods cacheController = nil
	if b.needsCacheInit {
		refresh := paramsProto.CacheRefreshDuration
		path, exists := os.LookupEnv("KUBECONFIG")
		if !exists {
			path = paramsProto.KubeconfigPath
		}
		controller, err := b.newCacheControllerFn(path, refresh, env)
		if err != nil {
			return nil, err
		}
		pods = controller
		env.ScheduleDaemon(func() { pods.Run(stopChan) })
		// ensure that any request is only handled after
		// a sync has occurred
		if env.Logger().VerbosityLevel(debugVerbosityLevel) {
			env.Logger().Infof("Waiting for kubernetes cache sync...")
		}
		if success := cache.WaitForCacheSync(stopChan, pods.HasSynced); !success {
			stopChan <- struct{}{}
			return nil, errors.New("cache sync failure")
		}
		if env.Logger().VerbosityLevel(debugVerbosityLevel) {
			env.Logger().Infof("Cache sync successful.")
		}
		b.needsCacheInit = false
	}
	kg := &handler{
		log:    env.Logger(),
		pods:   pods,
		params: *paramsProto,
	}
	return kg, nil
}

func (k *handler) Generate(inputs map[string]interface{}) (map[string]interface{}, error) {

	values := make(map[string]interface{})
	if id, found := serviceIdentifier(inputs, k.params.DestinationUidInputName, k.params.DestinationIpInputName); found && len(id) > 0 {
		k.addValues(values, id, k.params.DestinationPrefix, nil)
	}
	if k.skipIngressLookups(values) {
		return values, nil
	}
	if id, found := serviceIdentifier(inputs, k.params.SourceUidInputName, k.params.SourceIpInputName); found && len(id) > 0 {
		k.addValues(values, id, k.params.SourcePrefix, nil)
	}
	if id, found := serviceIdentifier(inputs, k.params.OriginUidInputName, k.params.OriginIpInputName); found && len(id) > 0 {
		k.addValues(values, id, k.params.OriginPrefix, nil)
	}
	return values, nil
}
func (k *handler) Generate2(inputs map[string]interface{}) (*adapter_template_kubernetes.Output, error) {

	values := make(map[string]interface{})

	var inst *adapter_template_kubernetes.Instance
	var out *adapter_template_kubernetes.Output

	if inst.DestinationUid != "" {
		k.addValues(values, inst.DestinationUid, k.params.DestinationPrefix, out)
	} else if inst.DestinationIp != nil {
		// TODO: update when support for golang net.IP is added to attribute.Bag
		var iface interface{} = inst.DestinationIp
		rawIP := iface.([]uint8)
		if len(rawIP) == net.IPv4len || len(rawIP) == net.IPv6len {
			ip := net.IP(rawIP)
			if !ip.IsUnspecified() {
				k.addValues(values, ip.String(), k.params.DestinationPrefix, out)
			}
		}
	}

	//if id, found := serviceIdentifier(inputs, k.params.DestinationUidInputName, k.params.DestinationIpInputName); found && len(id) > 0 {
	//	k.addValues(values, id, k.params.DestinationPrefix)
	//}
	//if k.skipIngressLookups(values) {
	//	return values, nil
	//}

	if k.skipIngressLookups2(out) {
		return out, nil
	}

	//if id, found := serviceIdentifier(inputs, k.params.SourceUidInputName, k.params.SourceIpInputName); found && len(id) > 0 {
	//	k.addValues(values, id, k.params.SourcePrefix)
	//}

	if inst.SourceUid != "" {
		k.addValues(values, inst.SourceUid, k.params.SourcePrefix, out)
	} else if inst.SourceIp != nil {
		// TODO: update when support for golang net.IP is added to attribute.Bag
		var iface interface{} = inst.SourceIp
		rawIP := iface.([]uint8)
		if len(rawIP) == net.IPv4len || len(rawIP) == net.IPv6len {
			ip := net.IP(rawIP)
			if !ip.IsUnspecified() {
				k.addValues(values, ip.String(), k.params.SourcePrefix, out)
			}
		}
	}

	//if id, found := serviceIdentifier(inputs, k.params.OriginUidInputName, k.params.OriginIpInputName); found && len(id) > 0 {
	//	k.addValues(values, id, k.params.OriginPrefix)
	//}
	if inst.OriginUid != "" {
		k.addValues(values, inst.OriginUid, k.params.OriginPrefix, out)
	} else if inst.OriginIp != nil {
		// TODO: update when support for golang net.IP is added to attribute.Bag
		var iface interface{} = inst.OriginIp
		rawIP := iface.([]uint8)
		if len(rawIP) == net.IPv4len || len(rawIP) == net.IPv6len {
			ip := net.IP(rawIP)
			if !ip.IsUnspecified() {
				k.addValues(values, ip.String(), k.params.OriginPrefix, out)
			}
		}
	}
	return out, nil
}

func (k *handler) addValues(vals map[string]interface{}, uid, valPrefix string, output *adapter_template_kubernetes.Output) {
	podKey := keyFromUID(uid)
	pod, found := k.pods.GetPod(podKey)
	if !found {
		if k.log.VerbosityLevel(debugVerbosityLevel) {
			k.log.Infof("could not find pod for (uid: %s, key: %s)", uid, podKey)
		}
		return
	}
	addPodValues(vals, valPrefix, k.params, pod, output)
}

func (k *handler) skipIngressLookups(values map[string]interface{}) bool {
	destSvcParam := k.params.DestinationPrefix + k.params.ServiceValueName
	return !k.params.LookupIngressSourceAndOriginValues && values[destSvcParam] == k.params.FullyQualifiedIstioIngressServiceName
}

func (k *handler) skipIngressLookups2(out *adapter_template_kubernetes.Output) bool {
	return !k.params.LookupIngressSourceAndOriginValues && out.DestinationService == k.params.FullyQualifiedIstioIngressServiceName
}

func newBuilder(cacheFactory controllerFactoryFn) *builder {
	return &builder{
		newCacheControllerFn: cacheFactory,
		needsCacheInit:       true,
		adapterConfig:        conf,
	}
	return &builder{}
}

// GetInfo returns the Info associated with this adapter implementation.
func GetInfo() adapter.Info {
	return adapter.Info{
		Name:        "KubernetesAttributeGenerator",
		Impl:        "istio.io/istio/mixer/adapter/kubernetes",
		Description: "Generate kubernetes attributes",
		SupportedTemplates: []string{
			adapter_template_kubernetes.TemplateName,
		},
		DefaultConfig: conf,

		NewBuilder: func() adapter.HandlerBuilder { return newBuilder(newCacheFromConfig) },
	}
}

func valueName(prefix, value string) string {
	return fmt.Sprintf("%s%s", prefix, value)
}

// name format examples that can be currently canonicalized:
//
// "hello:80",
// "hello",
// "hello.default:80",
// "hello.default",
// "hello.default.svc:80",
// "hello.default.svc",
// "hello.default.svc.cluster:80",
// "hello.default.svc.cluster",
// "hello.default.svc.cluster.local:80",
// "hello.default.svc.cluster.local",
func canonicalName(service, namespace, clusterDomain string) (string, error) {
	if len(service) == 0 {
		return "", errors.New("invalid service name: cannot be empty")
	}
	// remove any port suffixes (ex: ":80")
	splits := strings.SplitN(service, ":", 2)
	s := splits[0]
	if len(s) == 0 {
		return "", fmt.Errorf("invalid service name '%s': starts with ':'", service)
	}
	// error on ip addresses for now
	if ip := net.ParseIP(s); ip != nil {
		return "", errors.New("invalid service name: cannot canonicalize ip addresses at this time")
	}
	parts := strings.SplitN(s, ".", 3)
	if len(parts) == 1 {
		return parts[0] + "." + namespace + "." + clusterDomain, nil
	}
	if len(parts) == 2 {
		return s + "." + clusterDomain, nil
	}

	domParts := strings.Split(clusterDomain, ".")
	nameParts := strings.Split(parts[2], ".")

	if len(nameParts) >= len(domParts) {
		return s, nil
	}
	for i := len(nameParts); i < len(domParts); i++ {
		s = s + "." + domParts[i]
	}
	return s, nil
}

func serviceIdentifier(inputs map[string]interface{}, keys ...string) (string, bool) {
	for _, key := range keys {
		if id, found := inputs[key]; found {
			switch id.(type) {
			// TODO: update when support for golang net.IP is added to attribute.Bag
			case []uint8:
				rawIP := id.([]uint8)
				if len(rawIP) == net.IPv4len || len(rawIP) == net.IPv6len {
					ip := net.IP(rawIP)
					if !ip.IsUnspecified() {
						return ip.String(), true
					}
				}
			case string:
				return id.(string), true
			}
		}
	}
	return "", false
}

func keyFromUID(uid string) string {
	if ip := net.ParseIP(uid); ip != nil {
		return uid
	}
	fullname := strings.TrimPrefix(uid, kubePrefix)
	if strings.Contains(fullname, ".") {
		parts := strings.Split(fullname, ".")
		if len(parts) == 2 {
			return key(parts[1], parts[0])
		}
	}
	return fullname
}

func addPodValues(m map[string]interface{}, prefix string, params config.Params, p *v1.Pod, o *adapter_template_kubernetes.Output) {
	if o != nil {
		if prefix == "source" {
			if len(p.Labels) > 0 {
				o.SourceLabels = p.Labels
			}
			if len(p.Name) > 0 {
				//m[valueName(prefix, params.PodNameValueName)] = p.Name
			}
			if len(p.Namespace) > 0 {
				o.SourceNamespace = p.Namespace
			}
			if len(p.Spec.ServiceAccountName) > 0 {
				o.SourceServiceAccountName = p.Spec.ServiceAccountName
			}
			if len(p.Status.PodIP) > 0 {
				o.SourcePodIp = net.ParseIP(p.Status.PodIP)
			}
			if len(p.Status.HostIP) > 0 {
				//m[valueName(prefix, params.HostIpValueName)] = net.ParseIP(p.Status.HostIP)
			}
			if app, found := p.Labels[params.PodLabelForService]; found {
				n, err := canonicalName(app, p.Namespace, params.ClusterDomainName)
				if err == nil {
					o.SourceService = n
				}
			} else if app, found := p.Labels[params.PodLabelForIstioComponentService]; found {
				n, err := canonicalName(app, p.Namespace, params.ClusterDomainName)
				if err == nil {
					o.SourceService = n
				}
			}
		} else if prefix == "destination" {
			if len(p.Labels) > 0 {
				o.DestinationLabels = p.Labels
			}
			if len(p.Name) > 0 {
				//m[valueName(prefix, params.PodNameValueName)] = p.Name
			}
			if len(p.Namespace) > 0 {
				o.DestinationNamespace = p.Namespace
			}
			if len(p.Spec.ServiceAccountName) > 0 {
				o.DestinationServiceAccountName = p.Spec.ServiceAccountName
			}
			if len(p.Status.PodIP) > 0 {
				o.DestinationPodIp = net.ParseIP(p.Status.PodIP)
			}
			if len(p.Status.HostIP) > 0 {
				//m[valueName(prefix, params.HostIpValueName)] = net.ParseIP(p.Status.HostIP)
			}
			if app, found := p.Labels[params.PodLabelForService]; found {
				n, err := canonicalName(app, p.Namespace, params.ClusterDomainName)
				if err == nil {
					o.DestinationService = n
				}
			} else if app, found := p.Labels[params.PodLabelForIstioComponentService]; found {
				n, err := canonicalName(app, p.Namespace, params.ClusterDomainName)
				if err == nil {
					o.DestinationService = n
				}
			}

		} else if prefix == "origin" {

		}
	}
	if len(p.Labels) > 0 {
		m[valueName(prefix, params.LabelsValueName)] = p.Labels
	}
	if len(p.Name) > 0 {
		m[valueName(prefix, params.PodNameValueName)] = p.Name
	}
	if len(p.Namespace) > 0 {
		m[valueName(prefix, params.NamespaceValueName)] = p.Namespace
	}
	if len(p.Spec.ServiceAccountName) > 0 {
		m[valueName(prefix, params.ServiceAccountValueName)] = p.Spec.ServiceAccountName
	}
	if len(p.Status.PodIP) > 0 {
		m[valueName(prefix, params.PodIpValueName)] = net.ParseIP(p.Status.PodIP)
	}
	if len(p.Status.HostIP) > 0 {
		m[valueName(prefix, params.HostIpValueName)] = net.ParseIP(p.Status.HostIP)
	}
	if app, found := p.Labels[params.PodLabelForService]; found {
		n, err := canonicalName(app, p.Namespace, params.ClusterDomainName)
		if err == nil {
			m[valueName(prefix, params.ServiceValueName)] = n
		}
	} else if app, found := p.Labels[params.PodLabelForIstioComponentService]; found {
		n, err := canonicalName(app, p.Namespace, params.ClusterDomainName)
		if err == nil {
			m[valueName(prefix, params.ServiceValueName)] = n
		}
	}
}

const (
	// adapter vals
	name = "kubernetes"
	desc = "Provides platform specific functionality for the kubernetes environment"

	// parsing
	kubePrefix = "kubernetes://"

	// input/output naming
	sourceUID         = "sourceUID"
	destinationUID    = "destinationUID"
	originUID         = "originUID"
	sourceIP          = "sourceIP"
	destinationIP     = "destinationIP"
	originIP          = "originIP"
	sourcePrefix      = "source"
	destinationPrefix = "destination"
	originPrefix      = "origin"
	labelsVal         = "Labels"
	podNameVal        = "PodName"
	podIPVal          = "PodIP"
	hostIPVal         = "HostIP"
	namespaceVal      = "Namespace"
	serviceAccountVal = "ServiceAccountName"
	serviceVal        = "Service"

	// value extraction
	clusterDomain                      = "svc.cluster.local"
	podServiceLabel                    = "app"
	istioPodServiceLabel               = "istio"
	lookupIngressSourceAndOriginValues = false
	istioIngressSvc                    = "ingress.istio-system.svc.cluster.local"

	// cache invaliation
	// TODO: determine a reasonable default
	defaultRefreshPeriod = 5 * time.Minute
)

var (
	conf = &config.Params{
		KubeconfigPath:                        "",
		CacheRefreshDuration:                  defaultRefreshPeriod,
		SourceUidInputName:                    sourceUID,
		DestinationUidInputName:               destinationUID,
		OriginUidInputName:                    originUID,
		SourceIpInputName:                     sourceIP,
		DestinationIpInputName:                destinationIP,
		OriginIpInputName:                     originIP,
		ClusterDomainName:                     clusterDomain,
		PodLabelForService:                    podServiceLabel,
		PodLabelForIstioComponentService:      istioPodServiceLabel,
		SourcePrefix:                          sourcePrefix,
		DestinationPrefix:                     destinationPrefix,
		OriginPrefix:                          originPrefix,
		LabelsValueName:                       labelsVal,
		PodNameValueName:                      podNameVal,
		PodIpValueName:                        podIPVal,
		HostIpValueName:                       hostIPVal,
		NamespaceValueName:                    namespaceVal,
		ServiceAccountValueName:               serviceAccountVal,
		ServiceValueName:                      serviceVal,
		FullyQualifiedIstioIngressServiceName: istioIngressSvc,
		LookupIngressSourceAndOriginValues:    lookupIngressSourceAndOriginValues,
	}
)

func newCacheFromConfig(kubeconfigPath string, refreshDuration time.Duration, env adapter.Env) (cacheController, error) {
	if env.Logger().VerbosityLevel(debugVerbosityLevel) {
		env.Logger().Infof("getting kubeconfig from: %#v", kubeconfigPath)
	}
	config, err := getRESTConfig(kubeconfigPath)
	if err != nil || config == nil {
		return nil, fmt.Errorf("could not retrieve kubeconfig: %v", err)
	}
	if env.Logger().VerbosityLevel(debugVerbosityLevel) {
		env.Logger().Infof("getting k8s client from config")
	}
	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("could not create clientset for k8s: %v", err)
	}
	return newCacheController(clientset, refreshDuration, env), nil
}

func getRESTConfig(kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
}
