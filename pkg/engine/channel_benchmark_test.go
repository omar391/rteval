package engine

import (
	"context"
	"testing"
)

func BenchmarkChannelThroughput(b *testing.B) {
	ch := make(chan int, 1024*1024)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ch:
			}
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- i
	}
}
