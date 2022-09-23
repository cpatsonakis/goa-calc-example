#!/bin/bash

goa gen github.com/cpatsonakis/goa-calc-example/design/demo
sleep 2;
goa example github.com/cpatsonakis/goa-calc-example/design/demo 

rm -rf calc.go

echo '
package calcapi

import (
	"context"
	"fmt"
	"log"

	calc "github.com/cpatsonakis/goa-calc-example/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplicationPayload) (res string, err error) {
	s.logger.Print("calc.multiply")
	return fmt.Sprintf("%d", p.A*p.B), nil
}
' > calc.go

go build ./cmd/calc && go build ./cmd/calc-cli

./calc