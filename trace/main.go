package main

import (
	"context"
	"runtime/trace"
	"time"
)

func main() {
	Cappuccino("17")
}

func Cappuccino(orderID string) {
	ctx, task := trace.NewTask(context.Background(), "cappuccino")
	trace.Log(ctx, "orderID", orderID)

	milk := make(chan bool)
	espresso := make(chan bool)

	go func() {
		trace.WithRegion(ctx, "steamMilk", steamMilk)
		milk <- true
	}()

	go func() {
		trace.WithRegion(ctx, "extractCoffee", extractCoffee)
		espresso <- true
	}()

	go func() {
		defer task.End()
		<-espresso
		<-milk
		trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
	}()
}

func steamMilk() {
	time.Sleep(30 * time.Millisecond)
}

func extractCoffee() {
	time.Sleep(10 * time.Millisecond)
}

func mixMilkCoffee() {
	time.Sleep(5 * time.Millisecond)
}
