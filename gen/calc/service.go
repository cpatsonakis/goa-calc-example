// Code generated by goa v3.8.5, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen github.com/cpatsonakis/goa-calc-example/design/demo

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

// Describes the format and properties of an error returned due to a bad
// request.
type BadRequestError struct {
	Name      string
	Message   string
	OccuredAt string
}

// MultiplicationPayload is the payload type of the calc service multiply
// method.
type MultiplicationPayload struct {
	// First operand of multiplication payload
	A int
	// Second operand of multiplication payload
	B int
}

type StringError string

// Error returns an error description.
func (e *BadRequestError) Error() string {
	return "Describes the format and properties of an error returned due to a bad request."
}

// ErrorName returns "BadRequestError".
func (e *BadRequestError) ErrorName() string {
	return "bad_request"
}

// Error returns an error description.
func (e StringError) Error() string {
	return ""
}

// ErrorName returns "StringError".
func (e StringError) ErrorName() string {
	return "mul_error"
}