// Code generated by goa v3.10.1, DO NOT EDIT.
//
// calculator HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/cpatsonakis/goa-calc-example/design/calc-full -o
// calc-full

package server

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	calculator "github.com/cpatsonakis/goa-calc-example/calc-full/gen/calculator"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeAddResponse returns an encoder for responses returned by the
// calculator add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(string)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "text/plain")
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the calculator add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			a   int64
			b   int64
			err error

			params = mux.Vars(r)
		)
		{
			aRaw := params["a"]
			v, err2 := strconv.ParseInt(aRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("a", aRaw, "integer"))
			}
			a = v
		}
		{
			bRaw := params["b"]
			v, err2 := strconv.ParseInt(bRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("b", bRaw, "integer"))
			}
			b = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewAdditionPayload(a, b)

		return payload, nil
	}
}

// EncodeAddError returns an encoder for errors returned by the add calculator
// endpoint.
func EncodeAddError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "internal_server_error":
			var res *calculator.ErrorResultType
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			var res *calculator.ErrorResultType
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeSubtractResponse returns an encoder for responses returned by the
// calculator subtract endpoint.
func EncodeSubtractResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(string)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "text/plain")
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSubtractRequest returns a decoder for requests sent to the calculator
// subtract endpoint.
func DecodeSubtractRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			a   int64
			b   int64
			err error

			params = mux.Vars(r)
		)
		{
			aRaw := params["a"]
			v, err2 := strconv.ParseInt(aRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("a", aRaw, "integer"))
			}
			a = v
		}
		{
			bRaw := params["b"]
			v, err2 := strconv.ParseInt(bRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("b", bRaw, "integer"))
			}
			b = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewSubtractionPayload(a, b)

		return payload, nil
	}
}

// EncodeSubtractError returns an encoder for errors returned by the subtract
// calculator endpoint.
func EncodeSubtractError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "internal_server_error":
			var res *calculator.ErrorResultType
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSubtractInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			var res *calculator.ErrorResultType
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewSubtractBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeMultiplyResponse returns an encoder for responses returned by the
// calculator multiply endpoint.
func EncodeMultiplyResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(string)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "text/plain")
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeMultiplyRequest returns a decoder for requests sent to the calculator
// multiply endpoint.
func DecodeMultiplyRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			a   int64
			b   int64
			err error

			params = mux.Vars(r)
		)
		{
			aRaw := params["a"]
			v, err2 := strconv.ParseInt(aRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("a", aRaw, "integer"))
			}
			a = v
		}
		{
			bRaw := params["b"]
			v, err2 := strconv.ParseInt(bRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("b", bRaw, "integer"))
			}
			b = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewMultiplyMultiplicationPayload(a, b)

		return payload, nil
	}
}

// EncodeMultiplyError returns an encoder for errors returned by the multiply
// calculator endpoint.
func EncodeMultiplyError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "internal_server_error":
			var res *calculator.ErrorResultType
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewMultiplyInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			var res *calculator.ErrorResultType
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewMultiplyBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDivideResponse returns an encoder for responses returned by the
// calculator divide endpoint.
func EncodeDivideResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*calculator.DivisionResult)
		enc := encoder(ctx, w)
		body := NewDivideResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDivideRequest returns a decoder for requests sent to the calculator
// divide endpoint.
func DecodeDivideRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			a   int64
			b   int64
			err error

			params = mux.Vars(r)
		)
		{
			aRaw := params["a"]
			v, err2 := strconv.ParseInt(aRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("a", aRaw, "integer"))
			}
			a = v
		}
		{
			bRaw := params["b"]
			v, err2 := strconv.ParseInt(bRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("b", bRaw, "integer"))
			}
			b = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewDivideDivisionPayload(a, b)

		return payload, nil
	}
}

// EncodeDivideError returns an encoder for errors returned by the divide
// calculator endpoint.
func EncodeDivideError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "internal_server_error":
			var res *calculator.ErrorResultType
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDivideInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			var res *calculator.ErrorResultType
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDivideBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
