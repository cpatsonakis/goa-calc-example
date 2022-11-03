package openapiui

import (
	"context"
	"fmt"
	"net/http"

	goopenapi "github.com/go-openapi/runtime/middleware"
	goahttp "goa.design/goa/v3/http"
)

const swaggerUIEndpoint = "/swagger"

type swaggerUIServer struct {
	baseServer
}

func NewSwaggerUIServer(
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	serverExternalBaseURL string,
	openAPIFilepath string,
) (Server, error) {
	basesrv, err := newBaseServer(mux, decoder, encoder, serverExternalBaseURL, openAPIFilepath, swaggerUIEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error while creating swagger ui server: %w", err)
	}
	basesrv.mounts = append(basesrv.mounts, &MountPoint{
		Method:  "Swagger UI",
		Verb:    http.MethodGet,
		Pattern: swaggerUIEndpoint,
	})
	basesrv.openAPIUIHandler = newSwaggerUIHandler(mux, basesrv.openAPIFileEndpoint)
	return &swaggerUIServer{
		baseServer: *basesrv,
	}, nil
}

// Service returns the name of the service served.
func (s *swaggerUIServer) Service() string { return "swagger" }

// Use wraps the server handlers with the given middleware.
func (s *swaggerUIServer) Use(m func(http.Handler) http.Handler) {
	s.openAPIFileHandler = m(s.openAPIFileHandler)
	s.openAPIUIHandler = m(s.openAPIUIHandler)
}

// Mount configures the mux to serve the swagger endpoints.
func (s *swaggerUIServer) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

func (s *swaggerUIServer) GetMountPoints() []*MountPoint {
	return s.baseServer.mounts
}

func (s *swaggerUIServer) GetOpenAPIFileHandler() http.Handler {
	return s.baseServer.openAPIFileHandler
}

func (s *swaggerUIServer) GetOpenAPIUIHandler() http.Handler {
	return s.baseServer.openAPIUIHandler
}

func (s *swaggerUIServer) GetOpenAPIUIEndpoint() string {
	return swaggerUIEndpoint
}

func (s *swaggerUIServer) GetOpenAPIFileEndpoint() string {
	return s.baseServer.openAPIFileEndpoint
}

func newSwaggerUIHandler(mux goahttp.Muxer, openAPIFileURL string) http.Handler {
	swaggerUIOpts := goopenapi.SwaggerUIOpts{
		Path:    swaggerUIEndpoint,
		SpecURL: openAPIFileURL,
	}
	return goopenapi.SwaggerUI(swaggerUIOpts, mux)
}
