package main

import (
	"flag"

	"github.com/harlow/go-micro-services/services/rate"
	"github.com/harlow/go-micro-services/tracing"
)

func main() {
	var (
		port       = flag.String("port", "8080", "The server port")
		jaegerAddr = flag.String("jaegeraddr", "jaeger:6831", "Jaeger server addr")
	)
	flag.Parse()

	srv := &rate.Server{
		Tracer: tracing.Init("rate", *jaegerAddr),
		Port:   *port,
	}
	srv.Run()
}
