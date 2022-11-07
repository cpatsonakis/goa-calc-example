package design

import (
	. "goa.design/goa/v3/dsl"
)

var ErrorResultType = Type("ErrorResultType", func() {
	ErrorName("name", String, "Name of the error.", func() {
		Enum("bad_request", "internal_error")
	})
	Attribute("message", String, "Descriptive error message.", func() {
		Example("default", "Something went wrong.")
	})
	Attribute("occured_at", String, "Timestamp of error's occurence.", func() {
		Format(FormatDateTime)
	})
	Required("name", "message", "occured_at")
})
