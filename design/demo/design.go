package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for multiplying numbers, a Goa teaser")
	// Server("calc", func() {
	// 	// Host("localhost", func() {
	// 	// 	URI("http://localhost:8000")
	// 	// })
	// 	Host("default", func() {
	// 		URI("http://localhost:8000")
	// 	})
	// })
	cors.Origin("*", func() {
		cors.Methods("GET", "POST")
		cors.Headers("*")
		cors.Credentials()
		cors.MaxAge(600)
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

var BadRequestError = ResultType("application/json", func() {
	Description("Describes the format and properties of an error returned due to a bad request.")
	//Reference(ErrorResult)
	TypeName("BadRequestError")
	//ErrorName("name", String, "BadRequestError")
	Attribute("name", String, func() {
		Example("default", "kapoio name sfalmatos")
	})
	Attribute("message", String, func() {
		Example("default", "kapoio minima sfalmatos")
	})
	Attribute("occured_at", String, func() {
		Meta("openapi:example", "true")
		Format(FormatDateTime)
	})

	Required("name", "message", "occured_at")
})

var CustomErrorType = Type("CustomError", func() {
	// The "name" attribute is used to select the error response.
	// name should be set to either "internal_error" or "bad_request" by
	// the service method returning the error.
	Description("Describes the format and properties of an error returned due to a bad request.")
	ErrorName("name", String, "Name of error.")
	Attribute("message", String, "Message of error.")
	Attribute("occurred_at", String, "Time error occurred.", func() {
		Format(FormatDateTime)
	})
	Required("name", "message", "occurred_at")
})

// var CustomErrorTypeRef = Type("CustomErrorTypeRef", ErrorResult, func() {
// 	// The "name" attribute is used to select the error response.
// 	// name should be set to either "internal_error" or "bad_request" by
// 	// the service method returning the error.
// 	Description("Describes the format and properties of an error returned due to a bad request ref.")
// 	// ErrorName("name", String, "Name of error.")
// 	// Attribute("message", String, "Message of error.")
// 	// Attribute("occurred_at", String, "Time error occurred.", func() {
// 	// 	Format(FormatDateTime)
// 	// })
// 	ErrorResult.Find("temporary").UserExamples = []*expr.ExampleExpr{
// 		{Summary: "asdfasdf",
// 			Description: "52t24",
// 			Value:       false,
// 		},
// 	}
// 	View("default", func() {
// 		Attribute("name")
// 		Attribute("message")
// 		Attribute("id")
// 		Attribute("timeout", func() {
// 			Example("default", "false")
// 		})
// 		Attribute("fault", func() {
// 			Example("default", "false")
// 		})
// 	})

// 	Required("name", "message", "id", "temporary", "timeout", "fault")
// })

var CustomErrorTypeRef = Type("CustomErrorTypeRef", ErrorResult, func() {
	Meta("openapi:example", "false")
	Required("name", "message", "id", "temporary", "timeout", "fault")
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

		// Error("bad_request", StringError, func() {
		// 	Example("A stringified message describing why the request was rejected.")
		// })

		//Error("bad_request")
		// Error("bad_request", CustomErrorTypeRef, func() {
		// 	Meta("openapi:example", "false")
		// })

		Error("bad_request", BadRequestError)
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

			// Response("bad_request", StatusBadRequest, func() {
			// 	ContentType("text/plain")
			// })

			Response("bad_request", StatusBadRequest)

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
