package errorformat

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/cpatsonakis/goa-calc-example/goa-calc/gen/calc"
	goahttp "goa.design/goa/v3/http"
	"goa.design/goa/v3/middleware"
	goa "goa.design/goa/v3/pkg"
)

type errorResultTypeStatUser struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	OccuredAt string `json:"occured_at"`
}

func (e errorResultTypeStatUser) StatusCode() int {
	return http.StatusBadRequest
}

func ErrorFormatter(logger *log.Logger) func(ctx context.Context, err error) goahttp.Statuser {
	return func(ctx context.Context, err error) goahttp.Statuser {
		var (
			formattedError errorResultTypeStatUser
			requestId      string
		)
		formattedError.OccuredAt = time.Now().Format(time.RFC3339)
		requestId = ctx.Value(middleware.RequestIDKey).(string)
		log.Printf("[errorformat.ErrorFormatter] %s id=%s error=%#v\n", formattedError.OccuredAt, requestId, err)
		if errorResultType, ok := err.(*calc.ErrorResultType); ok {
			formattedError.Message = errorResultType.Message
			formattedError.Name = errorResultType.GoaErrorName()
		} else if errorResultType, ok := err.(*goa.ServiceError); ok {
			formattedError.Message = errorResultType.Error()
			formattedError.Name = BadRequestErrorName
		}
		return formattedError
	}
}
