package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("calculator", func() {
	Description("The calculator service performs legendary mathematical operations on numbers.")

	Docs(func() {
		Description("Specification")
		URL("http://there-is-no-documentation-for-calculator.com")
	})

	Error("internal_error", GenericHTTPError)
	Error("bad_request", GenericHTTPError)

	Method("add", func() {
		Description("Addition of two numbers.")

		Docs(func() {
			Description("Addition documentation.")
			URL("http://there-is-no-documentation-for-addition.com")
		})

		Payload(AdditionPayload, "First and second addition operands.", func() {
			Description("First and second addition operands.")
			Required("a", "b")
		})

		Result(Int64)

		HTTP(func() {
			GET("/add/{a}/{b}")
		})
	})

	Method("multiply", func() {
		Description("Multiplication of two numbers.")

		Docs(func() {
			Description("Multiplication documentation.")
			URL("http://there-is-no-documentation-for-multiplication.com")
		})

		Payload(MultiplicationPayload, "First and second multiplication operands.", func() {
			Description("First and second multiplication operands.")
			Required("a", "b")
		})

		Result(Int64)

		HTTP(func() {
			GET("/mul/{a}/{b}")
		})
	})
})
