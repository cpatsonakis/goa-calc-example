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

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("multiply", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand", func() {
				Example("default", 3)
			})
			Field(2, "b", Int, "Right operand", func() {
				Example("default", 5)
			})
			Required("a", "b")
		})

		Result(String, "Multiplication result.")

		HTTP(func() {
			GET("/multiply/{a}/{b}")
			Response(StatusOK, func() {
				ContentType("text/plain")
			})
		})
	})

	Files("/openapi3.json", "./gen/http/openapi3.json", func() {
		Description("JSON document containg the API's OpenAPI 3.0 definition.")
	})
	Files("/openapi.json", "./gen/http/openapi.json", func() {
		Description("JSON document containg the API's OpenAPI 2.0 definition.")
	})
})
