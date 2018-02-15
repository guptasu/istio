package main

import (

	"bytes"
	"context"
	"fmt"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/grpc"
	"io/ioutil"
	configpb "istio.io/api/mixer/v1/config"
	pbv "istio.io/api/mixer/v1/config/descriptor"
	rpc "istio.io/gogo-genproto/googleapis/google/rpc"
	grpcPkg "istio.io/istio/mixer/cmd/genericGRPC/pkg"
	"istio.io/istio/mixer/pkg/attribute"
	"istio.io/istio/mixer/pkg/expr"
	"istio.io/istio/mixer/pkg/il/compiled"
)

func main() {
	fds, err := getFileDescSet("/Users/guptasu/go/src/istio.io/istio/mixer/cmd/genericGRPC/descriptors/metric.pb")
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithCodec(ByteCodec{}),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reqBytes := getNewRequestBytes(
		`name: attr1`,
		map[string]interface{} {
			"attr1": "attr1Val",
		},
		fds,
	)
	rpcStatusBytes := getRpcStatusBytes()
	if err = conn.Invoke(
		context.Background(),
		"/mixer.adapter.metricentry.MetricEntryService/HandleMetricEntry",
		reqBytes,
		rpcStatusBytes); err != nil {
		panic(err)
	}

	result := rpc.Status{}
	gogoproto.Unmarshal(rpcStatusBytes.Bytes(), &result)
	fmt.Printf("remote adapter response: %v", result)
}

func getNewRequestBytes(instCfg string, attrs map[string]interface{}, fds *descriptor.FileDescriptorSet) *bytes.Buffer {
	var fd *descriptor.FileDescriptorProto
	fd = fds.File[3]
	instParamBytes, _ := grpcPkg.YamlToBytes(instCfg, fd, "InstanceParam")

	builder := compiled.NewBuilder(expr.NewFinder(manifest))
	assembler, _ := grpcPkg.NewAssemblerFor(fd, "Instance", "InstanceParam", instParamBytes, builder)

	buf := grpcPkg.GetBuffer()
	_ = assembler.Assemble(attribute.GetFakeMutableBagForTesting(attrs), buf)
	return bytes.NewBuffer(buf.Bytes())
}

func getRpcStatusBytes() *bytes.Buffer {
	r := rpc.Status{}
	bts2, _ := gogoproto.Marshal(&r)
	return bytes.NewBuffer(bts2)
}

func getFileDescSet(path string) (*descriptor.FileDescriptorSet, error) {
	byts, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	fds := &descriptor.FileDescriptorSet{}
	err = gogoproto.Unmarshal(byts, fds)

	return fds, err
}

var manifest = map[string]*configpb.AttributeManifest_AttributeInfo{
	"attr1": {
		ValueType: pbv.STRING,
	},
}

type ByteCodec struct{}

func (ByteCodec) Marshal(v interface{}) ([]byte, error) {
	return v.(*bytes.Buffer).Bytes(), nil
}

func (ByteCodec) Unmarshal(data []byte, v interface{}) error {
	_, err := v.(*bytes.Buffer).Write(data)
	return err
}

func (ByteCodec) String() string {
	return "byte buffer"
}
