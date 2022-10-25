package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	"goa.design/plugins/v3/cors"
)

// Server lists the docs service endpoint HTTP handlers.
type Server struct {
	Prefix              string
	ExternalURL         string
	OpenAPIJSONFilename string
	Mounts              []*MountPoint
	CORS                http.Handler
	StaticHandler       http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the swaggerui service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	prefix string,
	externalURL string,
	openAPIJSONFilename string,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Prefix:              prefix,
		ExternalURL:         externalURL,
		OpenAPIJSONFilename: openAPIJSONFilename,
		Mounts: []*MountPoint{
			// {"CORS", "OPTIONS", "/openapi3.json"},
			// {"CORS", "OPTIONS", "/openapi.json"},
			{"OpenAPI 3 JSON Mount Point", "GET", "/swaggerui/openapi3.json"},
			{"Static Assets Mount Point", "GET", "/swaggerui/{file}"},
		},
		//CORS: NewCORSHandler(),
	}

}

// Service returns the name of the service served.
func (s *Server) Service() string { return "swaggerui" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	//s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return []string{} }

// Mount configures the mux to serve the docs endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	//MountCORSHandler(mux, h.CORS)
	//fh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./static/swagger-ui/")))
	//fhJSON := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./static/openapi3.json")))
	//mux.Handle(http.MethodGet, "/swaggerui/openapi3.json", fhJSON.ServeHTTP)
	//mux.Handle(http.MethodGet, "/swaggerui/{file}", fh.ServeHTTP)
	//mux.Handle(http.MethodGet, "/swaggerui/{file}", goahttp.Replace("", "/swaggerui/", http.FileServer(http.Dir("./static/swagger-ui/"))).ServeHTTP)
} //func(ResponseWriter, *Request)

// Mount configures the mux to serve the docs endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountGoaCalcGenHTTPOpenapi3JSON configures the mux to serve GET request made
// to "/openapi3.json".
func MountGoaCalcGenHTTPOpenapi3JSON(mux goahttp.Muxer, h http.Handler) {

	mux.Handle("GET", "/openapi3.json", HandleDocsOrigin(h).ServeHTTP)
}

// MountGoaCalcGenHTTPOpenapiJSON configures the mux to serve GET request made
// to "/openapi.json".
func MountGoaCalcGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", HandleDocsOrigin(h).ServeHTTP)
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service docs.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleDocsOrigin(h)
	mux.Handle("OPTIONS", "/openapi3.json", h.ServeHTTP)
	mux.Handle("OPTIONS", "/openapi.json", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleDocsOrigin applies the CORS response headers corresponding to the
// origin for the service docs.
func HandleDocsOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
				w.Header().Set("Access-Control-Allow-Headers", "*")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
