GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
API_PROTO_FILES=$(shell find api -name forum.proto)
ERRORS_PROTO_FILES=$(shell find api -name errors.proto)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:./api \
	       --validate_out=paths=source_relative,lang=go:./api \
	       --go-grpc_out=paths=source_relative:./api \
	       --grpc-gateway_out=paths=source_relative:./api \
	       --openapi_out==paths=source_relative:. \
	       $(API_PROTO_FILES)

.PHONY: errors
# generate errors proto
errors:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:./api \
	       --go-errors_out=paths=source_relative:./api \
	       $(ERRORS_PROTO_FILES)

.PHONY: proto
# generate code from all proto files
proto:
	make api;
	make errors;

.PHONY: wire
# generate wire code
wire:
	go mod tidy
	go get -d github.com/google/wire/cmd/wire@latest
	go generate ./...

,PHONY: generate
generate:
	make proto;
	make wire;

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help