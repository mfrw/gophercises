package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/mfrw/gophercises/kit/hello/endpoints"
	"github.com/mfrw/gophercises/kit/hello/service"
	"github.com/oklog/run"
)

const (
	serviceName = "oc-gokit-example"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "svc", serviceName)
	}

	var handler http.Handler
	{
		svc := service.Service{}
		endpoints := endpoints.Endpoints{
			Hello: endpoints.MakeHelloEndpoint(svc),
		}

		var serverOptions []httptransport.ServerOption
		serverOptions = append(serverOptions, httptransport.ServerErrorLogger(logger))

		handler = svchttp.NewHTTPHandler(endpoints, serverOptions...)
	}

	var g run.Group
	// Start svc handler
	{
		var (
			listener, _ = net.Listen("tcp", ":0") // dynamic port assignment
			addr        = listener.Addr().String()
		)
		g.Add(func() error {
			logger.Log("msg", "service start", "transport", "http", "address", addr)
			return http.Serve(listener, handler)
		}, func(error) {
			listener.Close()
		})
	}
	// Setup signal handler
	{
		var (
			cancelInterrupt = make(chan struct{})
			c               = make(chan os.Signal, 2)
		)
		defer close(c)

		g.Add(func() error {
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("recived signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	logger.Log("exit", g.Run())
}
