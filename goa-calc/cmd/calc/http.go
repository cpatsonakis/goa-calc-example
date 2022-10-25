package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/cpatsonakis/goa-calc-example/goa-calc/errorformat"
	calc "github.com/cpatsonakis/goa-calc-example/goa-calc/gen/calc"
	calcsvr "github.com/cpatsonakis/goa-calc-example/goa-calc/gen/http/calc/server"
	docssvr "github.com/cpatsonakis/goa-calc-example/goa-calc/gen/http/docs/server"
	httpswagger "github.com/swaggo/http-swagger"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, calcEndpoints *calc.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		calcServer *calcsvr.Server
		docsServer *docssvr.Server
	)
	{
		eh := errorHandler(logger)
		ef := errorformat.ErrorFormatter(logger)
		calcServer = calcsvr.New(calcEndpoints, mux, dec, enc, eh, ef)
		docsServer = docssvr.New(nil, mux, dec, enc, eh, ef, nil, nil)
		if debug {
			servers := goahttp.Servers{
				calcServer,
				docsServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	calcsvr.Mount(mux, calcServer)
	docssvr.Mount(mux, docsServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}
	serveFile(mux)
	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler, ReadHeaderTimeout: time.Second * 60}

	for _, m := range calcServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range docsServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			logger.Printf("failed to shutdown: %v", err)
		}
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}

// func serveFile(mux goahttp.Muxer) {
// 	dir := http.Dir("static/swagger-ui")

// 	handler := http.StripPrefix("/swaggerui/", http.FileServer(dir))
// 	mux.Handle(http.MethodGet, "/swaggerui/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("%v\n", r)
// 		handler.ServeHTTP(w, r)
// 	})
// 	mux.Handle(http.MethodGet, "/swaggerui/{file}", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(r)
// 		handler.ServeHTTP(w, r)
// 	})
// }

func serveFile(mux goahttp.Muxer) {
	// httpswagger.Handler(
	// 	httpswagger.URL("http://localhost:8080/swagger/openapi3.json")
	// )
	mux.Handle(http.MethodGet, "/swagger/*", httpswagger.Handler(
		httpswagger.URL("http://localhost:8080/swagger/openapi3.json"),
	))
	dir := http.Dir("goa-calc/gen/http")

	handler := http.StripPrefix("/swagger/", http.FileServer(dir))
	mux.Handle(http.MethodGet, "/swagger/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v\n", r)
		handler.ServeHTTP(w, r)
	})
}
