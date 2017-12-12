package server

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
)

func Serve(processor thrift.TProcessor, port int32) error {
	fmt.Printf("starting Gateway server at 0.0.0.0:%d\n", port)

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	host := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	socket, err := thrift.NewTServerSocket(host)
	if err != nil {
		fmt.Printf("start Gateway server failed - %s", err)
		return err
	}

	fmt.Printf("started Gateway server at 0.0.0.0:%d\n", port)

	server := thrift.NewTSimpleServer4(processor, socket, transportFactory, protocolFactory)
	server.Serve()
	return nil
}
