package openapiui

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cpatsonakis/goa-calc-example/helpers"
	goahttp "goa.design/goa/v3/http"
)

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

type Server interface {
	Service() string
	Use(m func(http.Handler) http.Handler)
	Mount(mux goahttp.Muxer)
	GetMountPoints() []*MountPoint
	GetOpenAPIFileHandler() http.Handler
	GetOpenAPIUIHandler() http.Handler
	GetOpenAPIFileEndpoint() string
	GetOpenAPIUIEndpoint() string
}

func Mount(mux goahttp.Muxer, s Server) {
	mux.Handle(http.MethodGet, s.GetOpenAPIFileEndpoint(), s.GetOpenAPIFileHandler().ServeHTTP)
	mux.Handle(http.MethodGet, s.GetOpenAPIUIEndpoint(), s.GetOpenAPIUIHandler().ServeHTTP)
}

type baseServer struct {
	openAPIFileEndpoint string
	openAPIFileHandler  http.Handler
	openAPIUIHandler    http.Handler
	mounts              []*MountPoint
}

func newBaseServer(
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	serverExternalBaseURL string,
	openAPIFilepath string,
	uiEndpoint string,
) (*baseServer, error) {
	externalURL, err := validateExternalURLString(serverExternalBaseURL)
	if err != nil {
		return nil, fmt.Errorf("validateExternalURLString() returned error: %w", err)
	}
	if !helpers.FileExists(openAPIFilepath) {
		return nil, fmt.Errorf("OpenAPI file %s does not exist", openAPIFilepath)
	}
	openAPIFile, err := os.Open(openAPIFilepath)
	if err != nil {
		return nil, fmt.Errorf("os.Open() for OpenAPI file returned error: %w", err)
	}
	var openAPIFileContentMap map[string]interface{}
	err = json.NewDecoder(openAPIFile).Decode(&openAPIFileContentMap)
	if err != nil {
		return nil, fmt.Errorf("json.Decode() for OpenAPI file returned error: %w", err)
	}
	openAPIFileContentMap["servers"] = []map[string]string{
		{
			"url": externalURL.Scheme + "://" + externalURL.Host,
		},
	}
	openAPIFileByteBuffer := bytes.NewBuffer(nil)
	if err = json.NewEncoder(openAPIFileByteBuffer).Encode(openAPIFileContentMap); err != nil {
		return nil, fmt.Errorf("json.Encode() for modified OpenAPI file returned error: %w", err)
	}
	openAPIFileEndpoint := uiEndpoint + "/openapi3.json"
	return &baseServer{
		openAPIFileEndpoint: openAPIFileEndpoint,
		openAPIFileHandler:  newOpenAPIFileHandler(mux, openAPIFileEndpoint, openAPIFileByteBuffer),
		mounts: []*MountPoint{
			{"OpenAPI File", http.MethodGet, openAPIFileEndpoint},
		},
	}, nil
}

func newOpenAPIFileHandler(mux goahttp.Muxer, openAPIFileEndpoint string, openAPIFileByteBuffer *bytes.Buffer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Header().Add("Content-Length", strconv.FormatInt(int64(openAPIFileByteBuffer.Len()), 10))
		w.Write(openAPIFileByteBuffer.Bytes())
	})
}
