// Used code from: http://www.grpc.io/docs/quickstart/go.html
package listeners

import (
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/model"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type UnaryGrpcListener struct {
	port int
}

func NewUnaryGrpcListener(port int) *UnaryGrpcListener {
	return &UnaryGrpcListener{port: port}
}

type unaryGrpcServerReplier struct {
	q queues.Queue
}

func (l *UnaryGrpcListener) StartAccepting(q queues.Queue) {
	fmt.Printf("Starting gRPC listening on port %d\n", l.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", l.port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	r := unaryGrpcServerReplier{q: q}
	model.RegisterGrpcUnaryInputterServiceServer(s, r)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func (r unaryGrpcServerReplier) MakeRequest(ctx context.Context, in *model.Request) (*model.Empty, error) {
	r.q.Enqueue([]byte(in.Message))
	return &model.Empty{}, nil
}
