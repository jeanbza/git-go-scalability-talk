package benchmark

import (
	"context"
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/inputters"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"google.golang.org/grpc"
	"log"
	"sync"
	"testing"
    "google.golang.org/grpc/examples/helloworld/helloworld"
)

var g grpcListenerBenchmark = grpcListenerBenchmark{}

type grpcListenerBenchmark struct {
    l  *listeners.GrpcListener
    wg *sync.WaitGroup
    q  queues.Queue
    p  int
    c helloworld.GreeterClient
}

func BenchmarkGrpcListener(b *testing.B) {
	if g.l == nil {
		g.p = benchmark.GetOpenTcpPort()
		fmt.Println("Starting on port", g.p)

		g.wg = &sync.WaitGroup{}
		g.q = benchmark.NewWaitingQueue(g.wg)

		g.l = listeners.NewGrpcListener(g.p)
		go g.l.StartAccepting(g.q)

        conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", g.p), grpc.WithInsecure())
        if err != nil {
            log.Fatalf("did not connect: %v", err)
        }
        g.c = helloworld.NewGreeterClient(conn)
	}

	for i := 0; i < b.N; i++ {
		g.wg.Add(1)

        _, err := g.c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "world"})
        if err != nil {
            log.Fatalf("could not greet: %v", err)
        }
	}

	g.wg.Wait()
}