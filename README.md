# git-go-scalability-talk

## Installation

1. Install [protobuf](https://github.com/google/protobuf/releases)
1. `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`

## Running

## Regenerating protobufs

1. `protoc application/model/grpc_inputter.proto --go_out=plugins=grpc:.`