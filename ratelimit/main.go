package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	N         = 1000
	mod       = 30
	nrBurst   = 16
	nrWrokers = 4
	every     = time.Millisecond * 100
)

type Limiter struct {
	lim *rate.Limiter
}

func (l *Limiter) Run(fn func()) error {
	ctx := context.Background()
	err := l.lim.Wait(ctx)
	if err != nil {
		return err
	}

	fn()

	return nil
}

func main() {
	var wg sync.WaitGroup
	l := rate.NewLimiter(rate.Every(every), nrBurst)
	ctx := context.Background()
	workCh := make(chan int, nrBurst)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < N; i++ {
			workCh <- int(rand.Int31n(30))
		}
		close(workCh)
	}()

	for i := 0; i < nrWrokers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := range workCh {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					l.Wait(ctx)
					work(id, i%mod)
				}(i)
			}

		}(i)

	}

	wg.Wait()
}

func work(id, i int) {
	log.Printf("WorkerID: %3d Fib(%d) => %d\n", id, i, fib(i))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
