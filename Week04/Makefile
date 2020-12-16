GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%Z)
PROTO_FILES=$(shell find . -name *.proto)
GOOGLE_API=$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis


GOVARS += -X main.Version=$(VERSION) -X main.Branch=$(BRANCH) -X main.Revision=$(REVISION) -X main.BuildDate=$(BUILD_DATE)
LDFLAGS = -ldflags "-s -w  $(GOVARS)"

.PHONY: init
init:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: proto
proto:
	protoc -I . -I $(GOOGLE_API)\
		 --go_out=paths=source_relative:. \
	     --go-grpc_out=paths=source_relative:. \
		 --grpc-gateway_out=paths=source_relative:. \
         --grpc-gateway_opt logtostderr=true \
         --grpc-gateway_opt paths=source_relative \
         $(PROTO_FILES)

.PHONY: build
build:
	mkdir -p build
	go build $(LDFLAGS)  -o ./build/ ./...

.PHONY: test
test:
	go test -v ./... -cover
