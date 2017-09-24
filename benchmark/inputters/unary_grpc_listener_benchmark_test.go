package benchmark

import (
	"context"
	"github.com/jadekler/git-go-scalability-talk/application/model"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"testing"
)

func BenchmarkUnaryGrpcListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ug.wg.Add(1)

		_, err := ug.c.MakeRequest(context.Background(), &model.Request{Message: benchmark.VERY_LARGE_MESSAGE})
		if err != nil {
			panic(err)
		}
	}

	ug.wg.Wait()
}

func BenchmarkUnaryGrpcListenerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ug.wg.Add(1)

			_, err := ug.c.MakeRequest(context.Background(), &model.Request{Message: benchmark.SMALL_MESSAGE})
			if err != nil {
				panic(err)
			}
		}
	})

	ug.wg.Wait()
}
