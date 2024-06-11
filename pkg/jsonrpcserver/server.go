package jsonrpcserver

import "fmt"

type JSONRPCServer struct {
	storage *Storage
}

func NewServer(storage *Storage) *JSONRPCServer {
	return &JSONRPCServer{storage: storage}
}

func (s *JSONRPCServer) HandleRequest(word string, result *string) error {
	err := s.storage.Create(word)
	if err != nil {
		return fmt.Errorf("Error creating word: %v", err)
	}

	*result = "Received and stored the word: " + word
	return nil
}
