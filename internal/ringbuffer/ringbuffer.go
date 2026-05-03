package ringbuffer

import (
	"runtime"
	"sync/atomic"
)

// RingBuffer is a lock-free ring buffer implementation using atomic operations.
// It is designed for single-producer, single-consumer (SPSC) or multi-producer, single-consumer (MPSC) scenarios.
type RingBuffer[T any] struct {
	buffer []T
	size   uint64
	mask   uint64
	head   uint64 // write index
	tail   uint64 // read index
}

func New[T any](size uint64) *RingBuffer[T] {
	// Ensure size is a power of 2 for fast masking
	if size&(size-1) != 0 {
		panic("size must be a power of 2")
	}
	return &RingBuffer[T]{
		buffer: make([]T, size),
		size:   size,
		mask:   size - 1,
	}
}

func (rb *RingBuffer[T]) Put(val T) bool {
	for {
		head := atomic.LoadUint64(&rb.head)
		tail := atomic.LoadUint64(&rb.tail)
		if head-tail >= rb.size {
			return false // Buffer full
		}
		if atomic.CompareAndSwapUint64(&rb.head, head, head+1) {
			rb.buffer[head&rb.mask] = val
			return true
		}
		runtime.Gosched()
	}
}

func (rb *RingBuffer[T]) Get() (T, bool) {
	tail := atomic.LoadUint64(&rb.tail)
	head := atomic.LoadUint64(&rb.head)
	if tail == head {
		var zero T
		return zero, false // Buffer empty
	}
	val := rb.buffer[tail&rb.mask]
	atomic.StoreUint64(&rb.tail, tail+1)
	return val, true
}
