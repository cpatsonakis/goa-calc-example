package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	goahttp "goa.design/goa/v3/http"
	goahttpmiddle "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"

	"github.com/cpatsonakis/goa-calc-example/gen/calc"
	"github.com/cpatsonakis/goa-calc-example/gen/http/calc/server"
	kin "github.com/getkin/kin-openapi/openapi3"
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
	someFunc()
	swaggerUIStuff(mux)
	redocUIStuff(mux)
	dec := goahttp.RequestDecoder                        //# Set HTTP request decoder
	enc := goahttp.ResponseEncoder                       //# Set HTTP response encoder
	svr := server.New(endpoints, mux, dec, enc, eh, nil) //# Create Goa HTTP server
	server.Mount(mux, svr)                               //# Mount Goa server on mux
	//externalURL := "http://chpats.dlt.iti.gr:9000"
	httpsvr := &http.Server{ //# Create Go HTTP server
		Addr:    ":9000", //# Configure server address
		Handler: handler, //# Set request handler
	}
	fmt.Println("Paei na sikosei ton goa http server...")
	if err := httpsvr.ListenAndServe(); err != nil { //# Start HTTP server
		panic(err)
	}
}

func someFunc() {
	doc, err := kin.NewLoader().LoadFromFile("openapi3.json")
	if err != nil {
		panic(err)
	}
	serverURL, err := url.Parse("/")
	if err != nil {
		panic(err)
	}
	server, strArray, str := doc.Servers.MatchURL(serverURL)
	fmt.Printf("Server : %v\n", server)
	fmt.Printf("strArray : %v\n", strArray)
	fmt.Printf("str : %v\n", str)
	server.URL = "http://chpats.dlt.iti.gr:9000"
	bytesDoc, err := doc.MarshalJSON()
	if err != nil {
		panic(err)
	}
	newDocFile, err := os.OpenFile("./static/openapi3.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	newDocFile.Write(bytesDoc)
	newDocFile.Close()
}

func swaggerUIStuff(muxer goahttp.MiddlewareMuxer) {
	fh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./static/swagger-ui")))
	fhJSON := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./static/")))
	muxer.Handle(http.MethodGet,
		"/swaggerui/openapi3.json",
		fhJSON.ServeHTTP)
	muxer.Handle(http.MethodGet,
		"/swaggerui/{file}",
		fh.ServeHTTP)
	muxer.Handle(http.MethodGet,
		"/swaggerui/",
		fh.ServeHTTP)

}

func redocUIStuff(muxer goahttp.MiddlewareMuxer) {
	fh := http.StripPrefix("/redoc/", http.FileServer(http.Dir("./static/redoc")))
	fhJSON := http.StripPrefix("/redoc/", http.FileServer(http.Dir("./static/")))
	muxer.Handle(http.MethodGet,
		"/redoc/openapi3.json",
		fhJSON.ServeHTTP)
	muxer.Handle(http.MethodGet,
		"/redoc/{file}",
		fh.ServeHTTP)
	muxer.Handle(http.MethodGet,
		"/redoc/",
		fh.ServeHTTP)

}

// func swaggerUIStuff(muxer goahttp.MiddlewareMuxer) {
// 	// h := swgui.NewHandler("My Api", "/openapi3.json", "/")
// 	// swguiHandler := http.StripPrefix("/swaggerui/", h)
// 	swguiHandler := swgui.NewHandler("My Api", "/openapi3.json", "/swaggerui/")
// 	// muxer.Handle(http.MethodGet,
// 	// 	"/swaggerui/{file}",
// 	// 	swguiHandler.ServeHTTP)
// 	muxer.Handle(http.MethodGet,
// 		"/swaggerui/{file}",
// 		swguiHandler.ServeHTTP)
// 	muxer.Handle(http.MethodGet,
// 		"/swaggerui/",
// 		swguiHandler.ServeHTTP)
// 	// muxer.Handle(http.MethodGet,
// 	// 	"/swaggerui/",
// 	// 	http.StripPrefix("/swaggerui/", swguiHandler).ServeHTTP)
// }
