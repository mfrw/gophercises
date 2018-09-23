package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context, start int) <-chan int {
	dst := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dst)
				return
			case dst <- start:
				start++
			}
		}
	}()
	return dst
}

func main() {
	ctx, cancle := context.WithDeadline(context.Background(), time.Now().Add(300*time.Millisecond))
	defer cancle()

	var i int
	for v := range gen(ctx, 0) {
		i = v
	}
	fmt.Println(i)
}
