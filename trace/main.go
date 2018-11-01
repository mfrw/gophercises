package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
	"strconv"
	"sync/atomic"
	"time"
)

var orders uint64

func main() {
	http.HandleFunc("/coffee", makeCupa)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	var str string
	fmt.Print("Press Return to exit:")
	fmt.Fscanf(os.Stdin, "%s", &str)
}

func makeCupa(w http.ResponseWriter, r *http.Request) {
	current_order := atomic.AddUint64(&orders, 1)
	cos := strconv.FormatUint(current_order, 10)
	Cappuccino(cos)
	w.Header().Set("content-type", "text/plain")
	w.Write([]byte(cos))
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
