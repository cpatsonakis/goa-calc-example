package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("calculator", func() {
	Description("The calculator service performs legendary mathematical operations on integers.")

	Docs(func() {
		Description("Specification")
		URL("http://there-is-no-documentation-for-calculator.com")
	})

	HTTP(func() {
		Path("/calculator")

		Response("bad_request", StatusBadRequest, func() {
			ContentType("application/json")
		})

		Response("internal_server_error", StatusInternalServerError, func() {
			ContentType("application/json")
		})
	})

	Error("internal_server_error", ErrorResultType)
	Error("bad_request", ErrorResultType)

	Method("add", func() {
		Description("Addition of two integers.")
		Meta("openapi:summary", "Integer addition")

		Docs(func() {
			Description("Addition documentation.")
			URL("http://there-is-no-documentation-for-addition.com")
		})

		Payload(AdditionPayload)

		Result(String, "Addition result.", func() {
			Example("default", "8")
		})

		HTTP(func() {
			POST("/add/{a}/{b}")

			Response(func() {
				Description("Successful integer addition response.")
				Code(StatusOK)
				ContentType("text/plain")
			})

		})
	})

	Method("subtract", func() {
		Description("Subtraction of two numbers.")
		Meta("openapi:summary", "Integer subtraction")

		Docs(func() {
			Description("Subtraction documentation.")
			URL("http://there-is-no-documentation-for-subtraction.com")
		})

		Payload(SubtractionPayload)

		Result(String, "Subtraction result.", func() {
			Example("default", "2")
		})

		HTTP(func() {
			POST("/sub/{a}/{b}")

			Response(func() {
				Description("Successful integer subtraction response.")
				Code(StatusOK)
				ContentType("text/plain")
			})

		})
	})

	Method("multiply", func() {
		Description("Multiplication of two numbers.")
		Meta("openapi:summary", "Integer multiplication")

		Docs(func() {
			Description("Multiplication documentation.")
			URL("http://there-is-no-documentation-for-multiplication.com")
		})

		Payload(MultiplicationPayload)

		Result(String, "Multiplication result.", func() {
			Example("default", "15")
		})

		HTTP(func() {
			POST("/mul/{a}/{b}")

			Response(func() {
				Description("Successful integer multiplication response.")
				Code(StatusOK)
				ContentType("text/plain")
			})

		})
	})

	Method("divide", func() {
		Description("Division of two numbers.")
		Meta("openapi:summary", "Integer division")

		Docs(func() {
			Description("Division documentation.")
			URL("http://there-is-no-documentation-for-division.com")
		})

		Payload(DivisionPayload)

		Result(DivisionResult)

		HTTP(func() {
			POST("/div/{a}/{b}")

			Response(func() {
				Description("Successful integer division response.")
				Code(StatusOK)
			})

		})
	})
})
