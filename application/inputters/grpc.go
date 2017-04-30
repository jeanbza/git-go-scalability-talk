// Used code from: http://www.grpc.io/docs/quickstart/go.html
package listeners

import (
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/model"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
    "golang.org/x/net/context"
)

type GrpcListener struct {
	port int
}

func NewGrpcListener(port int) *GrpcListener {
	return &GrpcListener{port: port}
}

type grpcServerReplier struct {
	q queues.Queue
}

func (l *GrpcListener) StartAccepting(q queues.Queue) {
    fmt.Printf("Starting gRPC listening on port %d\n", l.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", l.port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	r := grpcServerReplier{q: q}
	model.RegisterGrpcInputterServiceServer(s, r)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
        panic(err)
	}
}

func (r grpcServerReplier) MakeRequest(ctx context.Context, in *model.Request) (*model.Empty, error) {
	r.q.Enqueue([]byte(in.Message))
	return &model.Empty{}, nil
}
