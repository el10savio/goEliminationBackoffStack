package stack

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewEBStack[int](2)

	// Push some values
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Pop values
	value := stack.Pop()
	assert.NotNil(t, value)
	assert.Equal(t, 3, *value)

	// Check if not empty
	assert.False(t, stack.IsEmpty())
}

func BenchmarkPushSingleThreaded(b *testing.B) {
	stack := NewEBStack[int](1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkPopSingleThreaded(b *testing.B) {
	stack := NewEBStack[int](1)
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkConcurrentPush(b *testing.B) {
	for _, numGoroutines := range []int{2, 4, 8, 16} {
		b.Run(fmt.Sprintf("goroutines-%d", numGoroutines), func(b *testing.B) {
			stack := NewEBStack[int](numGoroutines)
			wg := sync.WaitGroup{}
			itemsPerGoroutine := b.N / numGoroutines

			b.ResetTimer()

			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for j := 0; j < itemsPerGoroutine; j++ {
						stack.Push(j)
					}
				}()
			}
			wg.Wait()
		})
	}
}

func BenchmarkConcurrentPushPop(b *testing.B) {
	for _, numGoroutines := range []int{2, 4, 8, 16} {
		b.Run(fmt.Sprintf("goroutines-%d", numGoroutines), func(b *testing.B) {
			stack := NewEBStack[int](numGoroutines)
			wg := sync.WaitGroup{}
			opsPerGoroutine := b.N / numGoroutines

			b.ResetTimer()

			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for j := 0; j < opsPerGoroutine; j++ {
						if rand.Float32() < 0.5 {
							stack.Push(j)
						} else {
							stack.Pop()
						}
					}
				}()
			}
			wg.Wait()
		})
	}
}

func BenchmarkLatency(b *testing.B) {
	stack := NewEBStack[int](runtime.GOMAXPROCS(0))
	latencies := make([]time.Duration, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		start := time.Now()
		stack.Push(i)
		latencies[i] = time.Since(start)
	}

	// Calculate percentiles
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	b.ReportMetric(float64(latencies[len(latencies)*50/100]), "p50-ns")
	b.ReportMetric(float64(latencies[len(latencies)*90/100]), "p90-ns")
	b.ReportMetric(float64(latencies[len(latencies)*99/100]), "p99-ns")
}

func TestRaceConditions(t *testing.T) {
	stack := NewEBStack[int](runtime.GOMAXPROCS(0))
	numOps := 10000
	wg := sync.WaitGroup{}

	// Multiple producers
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				stack.Push(j)
			}
		}()
	}

	// Multiple consumers
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				stack.Pop()
			}
		}()
	}

	wg.Wait()
}

func TestThroughput(t *testing.T) {
	stack := NewEBStack[int](runtime.GOMAXPROCS(0))
	duration := 5 * time.Second
	numGoroutines := runtime.GOMAXPROCS(0)

	var ops atomic.Uint64
	done := make(chan struct{})

	// Start producer/consumer goroutines
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					stack.Push(1)
					stack.Pop()
					ops.Add(2)
				}
			}
		}()
	}

	// Run for specified duration
	time.Sleep(duration)
	close(done)

	opsPerSec := float64(ops.Load()) / duration.Seconds()
	t.Logf("Throughput: %.2f ops/sec with %d goroutines", opsPerSec, numGoroutines)
}
