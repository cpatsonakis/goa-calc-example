// Code generated by goa v3.8.5, DO NOT EDIT.
//
// docs endpoints
//
// Command:
// $ goa gen github.com/cpatsonakis/goa-calc-example/design/demo

package docs

import (
	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "docs" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "docs" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "docs" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}
