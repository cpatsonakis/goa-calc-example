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
	"github.com/gorilla/mux"
	goa "goa.design/goa/v3/pkg"
)

type svc struct{}

func (s *svc) Multiply(ctx context.Context, p *calc.MultiplicationPayload) (string, error) {
	return fmt.Sprintf("%d", p.A*p.B), nil
}

func mySillyErrorHandler(ctx context.Context, w http.ResponseWriter, err error) {
	// w.WriteHeader(http.StatusBadRequest)
	// bre := calc.BadRequestError {
	// 	Name: "kakouli request",
	// 	Message: err.Error(),
	// 	OccuredAt: time.Now().UTC().String(),
	// }
	// goahttp.ResponseEncoder(ctx)
	// w.Write(bre.)
	fmt.Println("MPAINEI STO XAZOULIKO ERROR HANDLER MOU")
}

type sillyStatUser struct {
	Name string `json:"name"`
}

func (s *sillyStatUser) StatusCode() int {
	return http.StatusBadRequest
}

func mySillyFormatter(err error) goahttp.Statuser {
	fmt.Println("MPAINEI STO XAZOULIKO FORMATTER MOU")
	fmt.Printf("Formatter Error: %s\n", err.Error())
	return &sillyStatUser{
		Name: "kati",
	}
	// return calc.BadRequestError{
	// 	Name:      "kati",
	// 	Message:   err.Error(),
	// 	OccuredAt: time.Now().UTC().String(),
	// }
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

func sillyEndpointMiddleware(ge goa.Endpoint) goa.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// fmt.Printf("To context mesa sto endpoint middleware einai:\n%-v\n", ctx)
		// fmt.Printf("To request einai:\n%v\n", request)
		// var ok bool
		// var mp *calc.MultiplicationPayload
		// mp, ok = request.(*calc.MultiplicationPayload)
		// if !ok {
		// 	err := fmt.Errorf("type casting error")
		// 	fmt.Println(err.Error())
		// 	return nil, err
		// }
		// fmt.Printf("To mp einai %d kai %d\n", mp.A, mp.B)
		// return mp.A * mp.B, nil
		response, err := ge(ctx, request)
		if err != nil {
			fmt.Printf("To error einai: %s\n", err.Error())
		}
		fmt.Printf("To response einai: %v\n", response)
		return response, err
	}
}

//errhandler func(context.Context, http.ResponseWriter, error)

func main() {
	eh := theErrorHandler(log.New(os.Stderr, "[cellar] ", log.Ltime))
	s := &svc{}                       //# Create Service
	endpoints := calc.NewEndpoints(s) //# Create endpoints
	//endpoints.Use(sillyEndpointMiddleware)
	//endpoints.Use(ErrorLogger(log.New(os.Stderr, "[newlogger] ", log.Ltime), ""))
	mux := goahttp.NewMuxer() //# Create HTTP muxer

	logger := log.New(os.Stdout, "[newlogger] ", log.Ltime)
	adapter := middleware.NewLogger(logger)
	var handler http.Handler = mux
	{
		handler = goahttpmiddle.Log(adapter)(handler)
		handler = goahttpmiddle.RequestID()(handler)
	}

	dec := goahttp.RequestDecoder                        //# Set HTTP request decoder
	enc := goahttp.ResponseEncoder                       //# Set HTTP response encoder
	svr := server.New(endpoints, mux, dec, enc, eh, nil) //# Create Goa HTTP server
	server.Mount(mux, svr)                               //# Mount Goa server on mux
	httpsvr := &http.Server{                             //# Create Go HTTP server
		Addr:    "localhost:8000", //# Configure server address
		Handler: handler,          //# Set request handler
	}
	go swaggerUIStuff()
	fmt.Println("Paei na sikosei ton goa http server...")
	if err := httpsvr.ListenAndServe(); err != nil { //# Start HTTP server
		panic(err)
	}
}

func swaggerUIStuff() {
	r := mux.NewRouter()

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./static/swagger-ui")))
	r.PathPrefix("/swaggerui/").Handler(sh).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
		// ... other code ...
	}
	fmt.Println("Paei na sikosei to swagger-ui server...")
	log.Fatal(srv.ListenAndServe())
}
