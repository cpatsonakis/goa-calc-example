package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for multiplying numbers, a Goa teaser")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var StringError = Type("StringError", String)

var MultiplicationPayload = Type("MultiplicationPayload", func() {
	Description("Type used by the multiply method containing both multiplication operands.")

	Attribute("a", Int, "First operand of multiplication payload", func() {
		Example("default", 3)
	})
	Attribute("b", Int, "Second operand of multiplication payload", func() {
		Example("default", 5)
	})

	Required("a", "b")
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Docs(func() {
		Description("Calculation service documentation:")
		URL("http://there-is-no-calculationn-docs.com")
	})

	Method("multiply", func() {

		Description("Multiply two integers a and b and get the result in the response's body.")
		Meta("openapi:summary", "Integer multiplication")

		Docs(func() {
			Description("Multiplication documentation")
			URL("http://there-is-no-multiplication-docs.com")
		})

		Payload(MultiplicationPayload)

		Error("mul_error", StringError, func() {
			Example("Server failure while attempting to multiply provided integer values.")
		})

		Error("bad_request", StringError, func() {
			Example("A stringified message describing why the request was rejected.")
		})

		Result(String, "Multiplication result.", func() {
			Example("default", "15")
		})

		HTTP(func() {
			GET("/multiply/{a}/{b}")

			Response(func() {
				Description("Successful integer multiplication response.")
				Code(StatusOK)
				ContentType("text/plain")
			})

			Response("mul_error", StatusInternalServerError, func() {
				Description("Returned when the service fails to calculate the outcome of an operation.")
				ContentType("text/plain")
			})

			Response("bad_request", StatusBadRequest, func() {
				ContentType("text/plain")
			})

			Response(StatusGatewayTimeout, func() {
				Body(Empty)
			})
		})
	})

	// Files("/openapi3.json", "./gen/http/openapi3.json", func() {
	// 	Description("JSON document containg the API's OpenAPI 3.0 definition.")
	// })
	// Files("/openapi.json", "./gen/http/openapi.json", func() {
	// 	Description("JSON document containg the API's OpenAPI 2.0 definition.")
	// })
})

var _ = Service("docs", func() {
	Description("The calc service performs operations on numbers.")

	Docs(func() {
		Description("Calculation service documentation:")
		URL("http://there-is-no-calculationn-docs.com")
	})

	Files("/openapi3.json", "./gen/http/openapi3.json", func() {
		Description("JSON document containg the API's OpenAPI 3.0 definition.")
	})
	Files("/openapi.json", "./gen/http/openapi.json", func() {
		Description("JSON document containg the API's OpenAPI 2.0 definition.")
	})
})
