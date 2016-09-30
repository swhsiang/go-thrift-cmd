package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"swhsiang/computing"
)

// SimpleHandler simple handler
type SimpleHandler struct{}

// NewHandler return initialized simpleHandler
func NewHandler() (*SimpleHandler, error) {
	return &SimpleHandler{}, nil
}

// Ping response the request from client
func (s *SimpleHandler) Ping() (*computing.StatusOfService, error) {
	fmt.Print("Pong!\n")
	return &computing.StatusOfService{Version: "0.0.1", Network: "9090"}, nil
}

// Compute base on input to compute
func (s *SimpleHandler) Compute(input *computing.InputOfComputing) (*computing.OutputOfComputing, error) {
	a, b := 11, 22
	return &computing.OutputOfComputing{Error: "Work well", Res: thrift.Int32Ptr(int32(a + b))}, nil
}

func main() {
	var transport thrift.TServerTransport
	var protocolFactory thrift.TProtocolFactory

	transport, _ = thrift.NewTServerSocket("localhost:9090")
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTBufferedTransportFactory(8192)
	handler, _ := NewHandler()
	processor := computing.NewComputingProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", 9090)
	_ = server.Serve()
}
