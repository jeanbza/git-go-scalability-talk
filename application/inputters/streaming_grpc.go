// Used code from: http://www.grpc.io/docs/quickstart/go.html
package listeners

import (
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/model"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"google.golang.org/grpc"
	"net"
)

type StreamingGrpcListener struct {
	port int
}

func NewStreamingGrpcListener(port int) *StreamingGrpcListener {
	return &StreamingGrpcListener{port: port}
}

type streamingGrpcServerReplier struct {
	q queues.Queue
}

func (l *StreamingGrpcListener) StartAccepting(q queues.Queue) {
	fmt.Printf("Starting gRPC listening on port %d\n", l.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", l.port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	model.RegisterGrpcStreamingInputterServiceServer(s, &streamingGrpcServerReplier{q: q})
	s.Serve(lis)
}

func (r streamingGrpcServerReplier) MakeRequest(request model.GrpcStreamingInputterService_MakeRequestServer) error {
	for {
		req, err := request.Recv()
		if err != nil {
			return err
		}

		r.q.Enqueue([]byte(req.Message))
	}

	return nil
}
