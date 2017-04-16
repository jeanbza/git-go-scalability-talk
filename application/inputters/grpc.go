// Used code from: http://www.grpc.io/docs/quickstart/go.html
package listeners

import (
    "golang.org/x/net/context"
    pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
    port = ":8000"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

//func main() {
//    lis, err := net.Listen("tcp", port)
//    if err != nil {
//        log.Fatalf("failed to listen: %v", err)
//    }
//    s := grpc.NewServer()
//    pb.RegisterGreeterServer(s, &server{})
//    // Register reflection service on gRPC server.
//    reflection.Register(s)
//    if err := s.Serve(lis); err != nil {
//        log.Fatalf("failed to serve: %v", err)
//    }
//}

