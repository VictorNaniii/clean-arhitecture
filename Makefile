include .env

LOCAL_BIN := $(CURDIR)/bin

.PHONY: install-deps get-deps generate generate-path generate-user-api

install-deps:
	set GOBIN=$(LOCAL_BIN)&& go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	set GOBIN=$(LOCAL_BIN)&& go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate: generate-path generate-user-api

generate-path:
	if not exist "pkg\user_v1" mkdir "pkg\user_v1"

generate-user-api:
	protoc --proto_path api\user_v1 --go_out=pkg\user_v1 --go_opt=paths=source_relative --go-grpc_out=pkg\user_v1 --go-grpc_opt=paths=source_relative api\user_v1\user.proto