package swagger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cpatsonakis/goa-calc-example/helpers"
	goopenapi "github.com/go-openapi/runtime/middleware"
	goahttp "goa.design/goa/v3/http"
)

const swaggerUIEndpoint = "/swagger"

type Server struct {
	swaggerFileEndpoint string
	swaggerFileHandler  http.Handler
	swaggerUIHandler    http.Handler
	Mounts              []*MountPoint
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

func New(
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	externalServerBaseUrl string,
	swaggerDocFilePath string,
) (*Server, error) {
	externalURL, err := validateExternalURLString(externalServerBaseUrl)
	if err != nil {
		return nil, fmt.Errorf("validateExternalURLString() returned error: %w", err)
	}
	if !helpers.FileExists(swaggerDocFilePath) {
		return nil, fmt.Errorf("swagger file %s does not exist", swaggerDocFilePath)
	}
	swaggerFile, err := os.Open(swaggerDocFilePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open() for swagger file returned error: %w", err)
	}
	var swaggerFileContentMap map[string]interface{}
	err = json.NewDecoder(swaggerFile).Decode(&swaggerFileContentMap)
	if err != nil {
		return nil, fmt.Errorf("json.Decode() for swagger file returned error: %w", err)
	}
	swaggerFileContentMap["servers"] = []map[string]string{
		{
			"url": externalURL.Scheme + "://" + externalURL.Host,
		},
	}
	byteBuffer := bytes.NewBuffer(nil)
	if err = json.NewEncoder(byteBuffer).Encode(swaggerFileContentMap); err != nil {
		return nil, fmt.Errorf("json.Encode() for modified swagger file returned error: %w", err)
	}
	swaggerFileEndpoint := swaggerUIEndpoint + "/openapi3.json"
	swaggerFileURL := externalServerBaseUrl + swaggerFileEndpoint
	return &Server{
		swaggerFileEndpoint: swaggerFileEndpoint,
		swaggerUIHandler:    newSwaggerUIHandler(mux, swaggerUIEndpoint, swaggerFileURL),
		swaggerFileHandler:  newSwaggerFileHandler(mux, swaggerFileEndpoint, byteBuffer),
		Mounts: []*MountPoint{
			{"Swagger File", http.MethodGet, swaggerFileEndpoint},
			{"Swagger UI", http.MethodGet, swaggerUIEndpoint},
		},
	}, nil
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "swagger" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.swaggerFileHandler = m(s.swaggerFileHandler)
	s.swaggerUIHandler = m(s.swaggerUIHandler)
}

// Mount configures the mux to serve the swagger endpoints.
func Mount(mux goahttp.Muxer, s *Server) {
	mux.Handle(http.MethodGet, s.swaggerFileEndpoint, s.swaggerFileHandler.ServeHTTP)
	mux.Handle(http.MethodGet, swaggerUIEndpoint, s.swaggerUIHandler.ServeHTTP)
}

// Mount configures the mux to serve the swagger endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

func newSwaggerUIHandler(mux goahttp.Muxer, swaggerUIEndpoint, swaggerFileURL string) http.Handler {
	swaggerUIOpts := goopenapi.SwaggerUIOpts{
		Path:    swaggerUIEndpoint,
		SpecURL: swaggerFileURL,
	}
	swaggerHandler := goopenapi.SwaggerUI(swaggerUIOpts, mux)
	return swaggerHandler
}

func newSwaggerFileHandler(mux goahttp.Muxer, swaggerFileEndpoint string, swaggerBytesBuffer *bytes.Buffer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Header().Add("Content-Length", strconv.FormatInt(int64(swaggerBytesBuffer.Len()), 10))
		w.Write(swaggerBytesBuffer.Bytes())
	})
}

func validateExternalURLString(urlString string) (*url.URL, error) {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, fmt.Errorf("url.Parse() returned error: %w", err)
	}
	if url.Scheme != "http" && url.Scheme != "https" {
		return nil, fmt.Errorf("invalid scheme %s specified, must be http or https", url.Scheme)
	}
	if !strings.HasPrefix(urlString, url.Scheme+"://"+url.Host) {
		return nil, fmt.Errorf("url string %s does not have a valid http(s) prefix, must be of the form <scheme>://<host>", urlString)
	}
	return url, nil
}
