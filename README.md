# git-go-scalability-talk

## Installation

1. Install [protoc](https://github.com/google/protobuf/releases) to compile protobufs
1. Install deps

    ```
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    go get github.com/gorilla/websocket
    go get golang.org/x/net/context
    go get google.golang.org/grpc
    ```

1. (optional) Regenerate protobuf clients+models `protoc application/model/*.proto --go_out=plugins=grpc:.`

## Running benchmarks

1. `go test ./... -bench .`

