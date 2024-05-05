installDeps:
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	# go install github.com/99designs/gqlgen

generate:
	rm -rf api-go/*
	
	rm -rf OUT
	mkdir OUT
	protoc -I proto --go_out ./OUT --go-grpc_out ./OUT --grpc-gateway_out ./OUT $(shell find proto -not -path "proto/google/*" -type f -iname '*.proto')
	mv OUT/go.leeeo.se/form-forge/api-go/* api-go
	rm -rf OUT

build:
	go mod tidy
	go build -race -o bin/${APP} ./cmd

.PHONY: lint
lint:
	golangci-lint run --fix --timeout=120s ./...

.PHONY: test
test:
	 go test ./...

test-coverage:
	go install github.com/ory/go-acc@latest
	go-acc -o coverage.out ./... -- -v
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

gql:
	cd gql && go run github.com/99designs/gqlgen generate
