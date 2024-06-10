.PHONY: run_http_server run_jsonrpc_server

run_http_server:
    go run cmd/httpserver/main.go

run_jsonrpc_server:
    go run cmd/jsonrpcserver/main.go
