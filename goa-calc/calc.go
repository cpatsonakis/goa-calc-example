package calcapi

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/cpatsonakis/goa-calc-example/goa-calc/errorformat"
	calc "github.com/cpatsonakis/goa-calc-example/goa-calc/gen/calc"
	"goa.design/goa/v3/middleware"
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

// Multiply two integers a and b and get the result in the response's body.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplicationPayload) (res string, err error) {
	var (
		requestId string
		ok        bool
	)
	if requestId, ok = ctx.Value(middleware.RequestIDKey).(string); !ok {
		return "", &calc.ErrorResultType{
			Name:      errorformat.InternalServerErrorName,
			Message:   "Failed to extract requestId from request context.",
			OccuredAt: time.Now().Format(time.RFC3339),
		}
	}
	s.logger.Printf("id=%s calc.multiply %v\n", requestId, *p)
	a := big.NewInt(p.A)
	b := big.NewInt(p.B)
	return a.Mul(a, b).String(), nil
}
