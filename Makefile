BINARY = auth-grpc

.PHONY: build
build:
	dep ensure
	protoc -I protos/ protos/auth.proto --go_out=plugins=grpc:protos
	go build -o ${BINARY}