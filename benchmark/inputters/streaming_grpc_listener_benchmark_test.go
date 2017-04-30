package benchmark

import (
	"context"
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/inputters"
	"github.com/jadekler/git-go-scalability-talk/application/model"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"google.golang.org/grpc"
	"sync"
	"testing"
)

var sg streamingGrpcListenerBenchmark = streamingGrpcListenerBenchmark{}

type streamingGrpcListenerBenchmark struct {
	l  *listeners.StreamingGrpcListener
	wg *sync.WaitGroup
	q  queues.Queue
	p  int
	s  model.GrpcStreamingInputterService_MakeRequestClient
}

func BenchmarkGrpcListener(b *testing.B) {
	if sg.l == nil {
		sg.p = benchmark.GetOpenTcpPort()

		sg.wg = &sync.WaitGroup{}
		sg.q = benchmark.NewWaitingQueue(sg.wg)

		sg.l = listeners.NewStreamingGrpcListener(sg.p)
		go sg.l.StartAccepting(sg.q)

		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", sg.p), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		client := model.NewGrpcStreamingInputterServiceClient(conn)
		stream, err := client.MakeRequest(context.Background())
		if err != nil {
			panic(err)
		}
		sg.s = stream
	}

	for i := 0; i < b.N; i++ {
		sg.wg.Add(1)

		err := sg.s.Send(&model.Request{Message: benchmark.VERY_LARGE_MESSAGE})
		if err != nil {
            panic(err)
		}
	}

	sg.wg.Wait()
}
