package openapiui

import (
	"context"
	"fmt"
	"net/http"

	goopenapi "github.com/go-openapi/runtime/middleware"
	goahttp "goa.design/goa/v3/http"
)

const redocUIEndpoint = "/redoc"

type redocUIServer struct {
	*baseServer
}

func NewRedocUIServer(
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	serverExternalBaseURL string,
	openAPIFilepath string,
) (*redocUIServer, error) {
	basesrv, err := newBaseServer(mux, decoder, encoder, serverExternalBaseURL, openAPIFilepath, redocUIEndpoint)
	if err != nil {
		return &redocUIServer{}, fmt.Errorf("error while creating redoc ui server: %w", err)
	}
	basesrv.mounts = append(basesrv.mounts, &MountPoint{
		Method:  "Redoc UI",
		Verb:    http.MethodGet,
		Pattern: redocUIEndpoint,
	})
	basesrv.openAPIUIHandler = newRedocUIHandler(mux, basesrv.openAPIFileEndpoint)
	return &redocUIServer{
		baseServer: basesrv,
	}, nil
}

// Service returns the name of the service served.
func (s redocUIServer) Service() string { return "redoc" }

// Use wraps the server handlers with the given middleware.
func (s *redocUIServer) Use(m func(http.Handler) http.Handler) {
	s.openAPIFileHandler = m(s.openAPIFileHandler)
	s.openAPIUIHandler = m(s.openAPIUIHandler)
}

// Mount configures the mux to serve the swagger endpoints.
func (s *redocUIServer) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

func (s redocUIServer) GetMountPoints() []*MountPoint {
	return s.baseServer.mounts
}

func (s redocUIServer) GetOpenAPIFileHandler() http.Handler {
	return s.baseServer.openAPIFileHandler
}

func (s redocUIServer) GetOpenAPIUIHandler() http.Handler {
	return s.baseServer.openAPIUIHandler
}

func (s redocUIServer) GetOpenAPIUIEndpoint() string {
	return redocUIEndpoint
}

func (s redocUIServer) GetOpenAPIFileEndpoint() string {
	return s.baseServer.openAPIFileEndpoint
}

func newRedocUIHandler(mux goahttp.Muxer, openAPIFileURL string) http.Handler {
	redocUIOpts := goopenapi.RedocOpts{
		Path:    redocUIEndpoint,
		SpecURL: openAPIFileURL,
	}
	return goopenapi.Redoc(redocUIOpts, mux)
}
