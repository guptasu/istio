package main

import (
	"log"

	"bytes"
	"fmt"
	"google.golang.org/grpc"
	"context"
	gogoproto "github.com/gogo/protobuf/proto"
	"io/ioutil"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"istio.io/istio/mixer/pkg/expr"
	"istio.io/istio/mixer/pkg/il/compiled"
	configpb "istio.io/api/mixer/v1/config"
	pbv "istio.io/api/mixer/v1/config/descriptor"
	"istio.io/istio/mixer/pkg/attribute"
	grpcPkg "istio.io/istio/mixer/cmd/genericGRPC/pkg"
	rpc "istio.io/gogo-genproto/googleapis/google/rpc"
)

const (
	port = ":50051"
	instCfg = `
name: as
`
)

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

type GRPCTunnelByteStream_RawBytesClient interface {
	Send(*bytes.Buffer) error
	Recv() (*bytes.Buffer, error)
	grpc.ClientStream

}

type gRPCTunnelByteStream_RawBytesClient struct {
	grpc.ClientStream
}

func (x *gRPCTunnelByteStream_RawBytesClient) Send(m *bytes.Buffer) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCTunnelByteStream_RawBytesClient) Recv() (*bytes.Buffer, error) {
	m := new(bytes.Buffer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithCodec(ByteCodec{}), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't dial: %v", err)
		panic(err)
	}
	defer conn.Close()

	method := "mixer.adapter.metricentry.MetricEntryService/HandleMetricEntry"

	fds, err := getFileDescSet("/Users/guptasu/go/src/istio.io/istio/mixer/cmd/genericGRPC/descriptors/metric.pb")
	if err != nil {
		panic(err)
	}
	//d := getRequestBytes()
	d2 := getNewRequestBytes(instCfg, fds)
	//r := getResponseBytes()
	r2 := getResponseBytes()
	err = conn.Invoke(context.Background(), method, d2, r2)
	if err != nil {
		panic(err)
	}

	fmt.Println(d2, r2)
	// validate
	//res := pb.HelloReply{}
	result := rpc.Status{}
	gogoproto.Unmarshal(r2.Bytes(), &result)
	fmt.Printf("result is %v", result)

}


func getNewRequestBytes(cfg string, fds *descriptor.FileDescriptorSet) *bytes.Buffer {
	var fd *descriptor.FileDescriptorProto
	fd = fds.File[3]
	instParamBytes, err := grpcPkg.YamlToBytes(cfg, fd,"InstanceParam")
	finder := expr.NewFinder(manifest)
	builder := compiled.NewBuilder(finder)
	if err != nil {
		panic(err)
	}
	// Create a new Assembler which, given an instance and instanceParam descriptor (and the bytes for the
	// instanceParam, creates an assembler that can create an instance.
	assembler, err := grpcPkg.NewAssemblerFor(fd, "Instance", "InstanceParam", instParamBytes, builder)
	if err != nil {
		panic(err)
	}

	attrs := map[string]interface{}{
		"as": "baz",
	}

	// Now with the assembler, try creating an instance directly into a proto buffer.
	bag := attribute.GetFakeMutableBagForTesting(attrs)
	buf := grpcPkg.GetBuffer()
	err = assembler.Assemble(bag, buf)
	if err != nil {
		panic(err)
	}

	//d := pb.HelloRequest{}
	//bts, _ := gogoproto.Marshal(&d)
	return bytes.NewBuffer(buf.Bytes())
}

func getResponseBytes() *bytes.Buffer {
	r :=rpc.Status{}
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
	"as": {
		ValueType: pbv.STRING,
	},
	"ai": {
		ValueType: pbv.INT64,
	},
}
