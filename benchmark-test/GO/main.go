package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func worker(start, end int, wg *sync.WaitGroup, result *float64, mu *sync.Mutex) {
	defer wg.Done()
	sum := 0.0
	for i := start; i < end; i++ {
		sum += math.Sqrt(float64(i))
	}
	mu.Lock()
	*result += sum
	mu.Unlock()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()

	const total = 100_000_000
	cores := runtime.NumCPU()
	chunk := total / cores

	var wg sync.WaitGroup
	var mu sync.Mutex
	var result float64

	for i := 0; i < cores; i++ {
		from := i * chunk
		to := from + chunk
		wg.Add(1)
		go worker(from, to, &wg, &result, &mu)
	}

	wg.Wait()
	fmt.Printf("Result: %.2f, Time: %.2f seconds\n", result, time.Since(start).Seconds())
}
