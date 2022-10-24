package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calculator", func() {
	Title("Legendary Calculator REST-based Service")
	Description("A completely legendary, innovative and ingenious web service that " +
		"provides REST-based calculator functionality.")
	Version("0.0.1")
	TermsOfService("http://there-are-no-terms-of-service.com")
	Contact(func() {
		Name("Christos Patsonakis")
		Email("cpatsonakis@iti.gr")
		URL("https://github.com/cpatsonakis")
	})
	License(func() {
		Name("None")
		URL("http://there-is-no-license.com")
	})
	Docs(func() {
		Description("Looking for documentation? Well, there isn't one!")
		URL("http://there-is-no-documentation.com")
	})
	Server("calculator-server", func() {
		Description("The calculator-server hosts the Legendary Calculator Service.")

		Services("calculator")

		// Host("production-only", func() {
		// 	Description("We only build production-oriented stuff.")

		// 	URI("http://{domain}/calc")

		// 	Variable("domain", String, "Domain Name", func() {
		// 		Default("localhost:8080")
		// 	})
		// })

		// Host("development", func() {
		// 	Description("Development server host.")

		// 	URI("http://localhost:8080/calc")
		// })
	})
})

var GenericHTTPError = Type("GenericHTTPError", func() {
	ErrorName("name", String, "String-encoded error name.")
	Attribute("code", Int64, "Integer-encoded error")
	Attribute("message", String, "Descriptive error message.")
	Attribute("occured_at", String, "Timestamp of error's occurence", func() {
		Format(FormatDateTime)
	})

	Required("name", "message", "occured_at")
})

var AdditionPayload = Type("AdditionPayload", func() {
	Description("Type used by the add method containing both addition operands.")

	Attribute("a", Int64, "First operand of addition payload")
	Attribute("b", Int64, "Second operand of addition payload")

	Required("a", "b")
})

var MultiplicationPayload = Type("MultiplicationPayload", func() {
	Description("Type used by the multiply method containing both multiplication operands.")

	Attribute("a", Int64, "First operand of multiplication payload")
	Attribute("b", Int64, "Second operand of multiplication payload")

	Required("a", "b")
})

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
