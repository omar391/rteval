package engine

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/omar391/rteval/internal/ringbuffer"
)

type Processor[T any] func(ctx context.Context, event T) error

type EventLoop[T any] struct {
	rb        *ringbuffer.RingBuffer[T]
	processor Processor[T]
}

func NewEventLoop[T any](size uint64, processor Processor[T]) *EventLoop[T] {
	return &EventLoop[T]{
		rb:        ringbuffer.New[T](size),
		processor: processor,
	}
}

func (el *EventLoop[T]) Emit(event T) bool {
	return el.rb.Put(event)
}

func (el *EventLoop[T]) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			event, ok := el.rb.Get()
			if !ok {
				continue
			}
			el.safeProcess(ctx, event)
		}
	}
}

func (el *EventLoop[T]) safeProcess(ctx context.Context, event T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic in event processor: %v\nStack trace:\n%s\n", r, debug.Stack())
		}
	}()

	if err := el.processor(ctx, event); err != nil {
		fmt.Printf("Error processing event: %v\n", err)
	}
}
