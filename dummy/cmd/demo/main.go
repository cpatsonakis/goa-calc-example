package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goahttpmiddle "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"

	"github.com/cpatsonakis/goa-calc-example/gen/calc"
	"github.com/cpatsonakis/goa-calc-example/gen/http/calc/server"
	goa "goa.design/goa/v3/pkg"
)

type svc struct{}

func (s *svc) Multiply(ctx context.Context, p *calc.MultiplicationPayload) (string, error) {
	return fmt.Sprintf("%d", p.A*p.B), nil
}

func theErrorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}

func ErrorLogger(l *log.Logger, prefix string) func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		// A Goa endpoint is itself a function.
		return goa.Endpoint(func(ctx context.Context, req interface{}) (interface{}, error) {
			// Call the original endpoint function.
			res, err := e(ctx, req)
			// Log any error.
			if err != nil {
				l.Printf("%s: %s", prefix, err.Error())
			}
			// Return endpoint results.
			return res, err
		})
	}
}

func main() {
	eh := theErrorHandler(log.New(os.Stderr, "[cellar] ", log.Ltime))
	s := &svc{}                       //# Create Service
	endpoints := calc.NewEndpoints(s) //# Create endpoints
	mux := goahttp.NewMuxer()         //# Create HTTP muxer
	logger := log.New(os.Stdout, "[newlogger] ", log.Ltime)
	adapter := middleware.NewLogger(logger)
	var handler http.Handler = mux
	{
		handler = goahttpmiddle.Log(adapter)(handler)
		handler = goahttpmiddle.RequestID()(handler)
	}
	swaggerUIStuff(mux)
	dec := goahttp.RequestDecoder                        //# Set HTTP request decoder
	enc := goahttp.ResponseEncoder                       //# Set HTTP response encoder
	svr := server.New(endpoints, mux, dec, enc, eh, nil) //# Create Goa HTTP server
	server.Mount(mux, svr)                               //# Mount Goa server on mux
	httpsvr := &http.Server{                             //# Create Go HTTP server
		Addr:    "localhost:9000", //# Configure server address
		Handler: handler,          //# Set request handler
	}
	fmt.Println("Paei na sikosei ton goa http server...")
	if err := httpsvr.ListenAndServe(); err != nil { //# Start HTTP server
		panic(err)
	}
}

func swaggerUIStuff(muxer goahttp.MiddlewareMuxer) {
	fh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./static/swagger-ui")))
	muxer.Handle(http.MethodGet,
		"/swaggerui/{file}",
		fh.ServeHTTP)
	muxer.Handle(http.MethodGet,
		"/swaggerui/",
		fh.ServeHTTP)

}
