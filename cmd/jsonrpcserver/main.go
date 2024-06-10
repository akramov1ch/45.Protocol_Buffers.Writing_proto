package main

import (
	"45HW/pkg/config"
	"45HW/pkg/jsonrpcserver"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	config.LoadConfig()
	storage := jsonrpcserver.NewStorage()
	server := jsonrpcserver.NewServer(storage)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":"+config.Conf.JSONRPCServerPort)
	if err != nil {
		log.Fatal("Error listening:", err)
	}
	log.Printf("Serving JSON-RPC server on port %s", config.Conf.JSONRPCServerPort)
	log.Fatal(http.Serve(l, nil))
}
