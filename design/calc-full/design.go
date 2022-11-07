package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
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
	cors.Origin("*", func() {
		cors.Methods("GET", "POST")
		cors.Headers("*")
		cors.Credentials()
		cors.MaxAge(600)
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
