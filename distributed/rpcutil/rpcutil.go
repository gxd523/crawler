package rpcutil

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func NewServeRpc(host string, service interface{}) error {
	_ = rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
