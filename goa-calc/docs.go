package calcapi

import (
	"log"

	docs "github.com/cpatsonakis/goa-calc-example/goa-calc/gen/docs"
)

// docs service example implementation.
// The example methods log the requests and return zero values.
type docssrvc struct {
	logger *log.Logger
}

// NewDocs returns the docs service implementation.
func NewDocs(logger *log.Logger) docs.Service {
	return &docssrvc{logger}
}
