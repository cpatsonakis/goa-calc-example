package main

import (
	"context"
	"encoding/json"
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
	goopenapi "github.com/go-openapi/runtime/middleware"
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
	//serveSwaggerUI(mux)
	serveSwaggerUI("/home/kasdeya/repos/git/cpatsonakis/goa-calc-example/goa-calc/gen/http/openapi3.json",
		"http://localhost:8080", mux)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}
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

// func serveSwaggerUI(mux goahttp.Muxer) {
// 	dirString, err := os.Getwd()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Current wd: %s\n", dirString)

// 	execPath, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Exec path: %s\n", execPath)

// 	dir := http.Dir("goa-calc/swagger")

// 	handler := http.StripPrefix("/swagger/", http.FileServer(dir))
// 	mux.Handle(http.MethodGet, "/swagger/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("%v\n", r)
// 		handler.ServeHTTP(w, r)
// 	})
// 	mux.Handle(http.MethodGet, "/swagger/{file}", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(r)
// 		handler.ServeHTTP(w, r)
// 	})
// }

// FileExists returns a bool indicating whether a file exists or not. Also
// in the case the input filename is a directory, even if it "exists", the function
// will also return false as it is not a file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

var bufferedSwaggerFile map[string]interface{}

func serveSwaggerUI(swaggerFilePath, externalURL string, mux goahttp.Muxer) {
	if !FileExists(swaggerFilePath) {
		fmt.Printf("Swagger file %s does not exist.\n", swaggerFilePath)
		return
	}
	swaggerFile, err := os.Open(swaggerFilePath)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(swaggerFile).Decode(&bufferedSwaggerFile)
	if err != nil {
		panic(err)
	}
	//,"servers":[{"url":"http://localhost:80","description":"Default server for calc"
	bufferedSwaggerFile["servers"] = []map[string]string{
		{
			"url": externalURL,
		},
	}

	mux.Handle(http.MethodGet, "/swagger/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MPIKE STIN SERVE FILE TOU OPENAPI3.JSON")
		err = json.NewEncoder(w).Encode(bufferedSwaggerFile)
		if err != nil {
			panic(err)
		}
		//http.ServeFile(w, r, swaggerFilePath)
	})

	swaggerUIOpts := goopenapi.SwaggerUIOpts{
		Path:    "/swagger",
		SpecURL: externalURL + "/swagger/openapi3.json",
	}
	swaggerHandler := goopenapi.SwaggerUI(swaggerUIOpts, mux)
	mux.Handle(http.MethodGet, "/swagger", swaggerHandler.ServeHTTP)
}
