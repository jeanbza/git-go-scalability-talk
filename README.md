# git-go-scalability-talk


## Installation

1. Install [protobuf](https://github.com/google/protobuf/releases)
1. `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`

## Running benchmarks

1. `cd benchmark && go test ./... -bench .`

## Regenerating protobufs

1. `protoc application/model/*.proto --go_out=plugins=grpc:.`
