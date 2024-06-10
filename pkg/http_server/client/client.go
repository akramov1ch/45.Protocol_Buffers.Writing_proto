package client

import (
	"45HW/pkg/config"
	"net/rpc"
)

func SendRequest(word string) (string, error) {
	client, err := rpc.DialHTTP("tcp", "localhost:"+config.Conf.JSONRPCServerPort)
	if err != nil {
		return "", err
	}
	var result string
	err = client.Call("JSONRPCServer.HandleRequest", word, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}
