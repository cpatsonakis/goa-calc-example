package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for multiplying numbers - Goa's teaser (with some minor additions for experimentation purposes).")
	cors.Origin("*", func() {
		cors.Methods("GET", "POST")
		cors.Headers("*")
		cors.Credentials()
		cors.MaxAge(600)
	})
})

var MultiplicationPayload = Type("MultiplicationPayload", func() {
	Description("Type used by the multiply method containing both multiplication operands.")

	Attribute("a", Int64, "First operand of the multiplication operation.", func() {
		Example("default", 3)
	})
	Attribute("b", Int64, "Second operand of the multiplication operation.", func() {
		Example("default", 5)
	})

	Required("a", "b")
})

var ErrorResultType = Type("ErrorResultType", func() {
	//Description("Generic error type returned by the service.")
	ErrorName("name", String, "Name of the error.")
	Attribute("message", String, "Descriptive error message.", func() {
		Example("default", "Something went wrong.")
	})
	Attribute("occured_at", String, "Timestamp of error's occurence.", func() {
		Format(FormatDateTime)
	})
	Required("name", "message", "occured_at")
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Docs(func() {
		Description("Calculation service documentation:")
		URL("http://there-is-no-calculationn-docs.com")
	})

	Error("bad_request", ErrorResultType, func() {
		Description("Bad request response.")
	})
	Error("internal_error", ErrorResultType)

	Method("multiply", func() {

		Description("Multiply two integers a and b and get the result in the response's body.")
		Meta("openapi:summary", "Integer multiplication")

		Docs(func() {
			Description("Multiplication documentation")
			URL("http://there-is-no-multiplication-docs.com")
		})

		Payload(MultiplicationPayload)

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

			Response("internal_error", StatusInternalServerError, func() {
				Description("Internal server error response.")
			})

			Response("bad_request", StatusBadRequest, func() {
				Description("Bad request server response.")
			})
		})
	})

})

var _ = Service("docs", func() {
	Description("The docs service allows the retrieval of OpenAPI-compliant documentation.")

	Files("/openapi3.json", "./goa-calc/gen/http/openapi3.json", func() {
		Meta("openapi:summary", "OpenAPI 3.0 service definition.")
		Description("Retrieve a JSON document containing the API's OpenAPI 3.0 definition.")
	})

	Files("/openapi.json", "./goa-calc/gen/http/openapi.json", func() {
		Meta("openapi:summary", "OpenAPI 2.0 service definition.")
		Description("Retrieve a JSON document containing the API's OpenAPI 2.0 definition.")
	})
})
