package main

import (
	"context"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Dice struct {
	facesCount int
}

func NewDice(facesCount int) Dice {
	rand.Seed(time.Now().UnixNano())
	return Dice{facesCount: facesCount}
}

func (d Dice) RollOnce() int {
	result := rand.Intn(d.facesCount) + 1
	return result
}

func (d Dice) RollMany(n int) int {
	var result int64 = 0

	handler := func(ctx context.Context, input <-chan int, wg *sync.WaitGroup) {
		for {
			select {
			case <-input:
				roll := d.RollOnce()
				atomic.AddInt64(&result, int64(roll))
				wg.Done()
			case <-ctx.Done():
				return
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	input := make(chan int)
	defer close(input)
	wg := &sync.WaitGroup{}

	for i := 0; i < runtime.NumCPU(); i++ {
		go handler(ctx, input, wg)
	}

	wg.Add(n)

	for i := 0; i < int(n); i++ {
		input <- i
	}

	wg.Wait()

	cancel()

	return int(result)
}
