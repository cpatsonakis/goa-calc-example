package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("calculator", func() {
	Title("Legendary Integer Calculator Service")
	Description("A completely legendary, innovative and ingenious web service that " +
		"provides REST-based, integer calculator functionality.")
	Version("0.0.1")
	TermsOfService("http://there-are-no-terms-of-service.com")
	Contact(func() {
		Name("Christos Patsonakis")
		Email("cpatsonakis@iti.gr")
		URL("https://github.com/cpatsonakis")
	})
	License(func() {
		Name("None License")
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
})
