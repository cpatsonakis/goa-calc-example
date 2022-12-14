// Code generated by goa v3.10.1, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen github.com/cpatsonakis/goa-calc-example/design/goa-calc -o goa-calc

package calc

import (
	"context"
)

// The calc service performs operations on numbers.
type Service interface {
	// Multiply two integers a and b and get the result in the response's body.
	Multiply(context.Context, *MultiplicationPayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"multiply"}

type ErrorResultType struct {
	// Name of the error.
	Name string
	// Descriptive error message.
	Message string
	// Timestamp of error's occurence.
	OccuredAt string
}

// MultiplicationPayload is the payload type of the calc service multiply
// method.
type MultiplicationPayload struct {
	// First operand of the multiplication operation.
	A int64
	// Second operand of the multiplication operation.
	B int64
}

// Error returns an error description.
func (e *ErrorResultType) Error() string {
	return ""
}

// ErrorName returns "ErrorResultType".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *ErrorResultType) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "ErrorResultType".
func (e *ErrorResultType) GoaErrorName() string {
	return e.Name
}
