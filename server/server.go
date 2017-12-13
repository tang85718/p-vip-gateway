package server

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
)

func Serve(processor thrift.TProcessor, port string) error {


	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	host := fmt.Sprintf("0.0.0.0:%s", port)
	socket, err := thrift.NewTServerSocket(host)
	if err != nil {
		fmt.Printf("start Gateway server failed - %s", err)
		return err
	}

	fmt.Printf("Listening at %s", host)

	server := thrift.NewTSimpleServer4(processor, socket, transportFactory, protocolFactory)
	server.Serve()
	return nil
}
