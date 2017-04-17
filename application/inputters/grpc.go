// Used code from: http://www.grpc.io/docs/quickstart/go.html
package listeners

import (
    "golang.org/x/net/context"
    pb "google.golang.org/grpc/examples/helloworld/helloworld"
    "github.com/jadekler/git-go-scalability-talk/application/queues"
    "net"
    "log"
    "google.golang.org/grpc/reflection"
    "google.golang.org/grpc"
    "fmt"
)

type GrpcListener struct {
    port int
}

func NewGrpcListener(port int) *GrpcListener {
    return &GrpcListener{port: port}
}

func (l *GrpcListener) StartAccepting(q queues.Queue) {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", l.port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, l)

    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func (s *GrpcListener) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}