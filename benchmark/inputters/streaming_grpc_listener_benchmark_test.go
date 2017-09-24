package benchmark

import (
	"context"
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/model"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"google.golang.org/grpc"
	"testing"
)

func BenchmarkStreamingGrpcListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sg.wg.Add(1)

		err := sg.s.Send(&model.Request{Message: benchmark.SMALL_MESSAGE})
		if err != nil {
			panic(err)
		}
	}

	sg.wg.Wait()
}

func BenchmarkStreamingGrpcListenerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", sg.p), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		client := model.NewGrpcStreamingInputterServiceClient(conn)

		stream, err := client.MakeRequest(context.Background())
		if err != nil {
			panic(err)
		}
		s := stream

		for pb.Next() {
            sg.wg.Add(1)
			err := s.Send(&model.Request{Message: benchmark.SMALL_MESSAGE})
			if err != nil {
				panic(err)
			}
		}
	})

	sg.wg.Wait()
}
