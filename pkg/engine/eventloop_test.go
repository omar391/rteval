package engine

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestEventLoop(t *testing.T) {
	var wg sync.WaitGroup
	processedCount := 0
	mu := sync.Mutex{}

	processor := func(ctx context.Context, event int) error {
		mu.Lock()
		processedCount++
		mu.Unlock()
		wg.Done()
		return nil
	}

	el := NewEventLoop[int](1024, processor)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		_ = el.Run(ctx)
	}()

	numEvents := 100
	wg.Add(numEvents)
	for i := 0; i < numEvents; i++ {
		if !el.Emit(i) {
			t.Errorf("Failed to emit event %d", i)
		}
	}

	wg.Wait()
	if processedCount != numEvents {
		t.Errorf("Expected %d processed events, got %d", numEvents, processedCount)
	}
}

func BenchmarkEventLoopThroughput(b *testing.B) {
	processor := func(ctx context.Context, event int) error {
		return nil
	}

	el := NewEventLoop[int](1024*1024, processor)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_ = el.Run(ctx)
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		el.Emit(i)
	}
}
