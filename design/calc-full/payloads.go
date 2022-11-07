package design

import (
	. "goa.design/goa/v3/dsl"
)

var AdditionPayload = Type("AdditionPayload", func() {
	Description("Type used by the add method containing both addition operands.")

	Attribute("a", Int64, "First operand of addition payload", func() {
		Example("default", 3)
	})
	Attribute("b", Int64, "Second operand of addition payload", func() {
		Example("default", 5)
	})

	Required("a", "b")
})

var MultiplicationPayload = Type("MultiplicationPayload", func() {
	Description("Type used by the multiply method containing both multiplication operands.")

	Attribute("a", Int64, "First operand of multiplication payload", func() {
		Example("default", 3)
	})
	Attribute("b", Int64, "Second operand of multiplication payload", func() {
		Example("default", 5)
	})

	Required("a", "b")
})

var SubtractionPayload = Type("SubtractionPayload", func() {
	Description("Type used by the subtract method containing both subtraction operands.")

	Attribute("a", Int64, "First operand of subtraction payload", func() {
		Example("default", 5)
	})
	Attribute("b", Int64, "Second operand of subtraction payload", func() {
		Example("default", 3)
	})

	Required("a", "b")
})

var DivisionPayload = Type("DivisionPayload", func() {
	Description("Type used by the divide method containing both division operands.")

	Attribute("a", Int64, "First operand (nominator) of division payload", func() {
		Example("default", 8)
	})
	Attribute("b", Int64, "Second operand (denominator) of division payload", func() {
		Example("default", 2)
	})

	Required("a", "b")
})

var DivisionResult = Type("DivisionResult", func() {
	Description("Type used")

	Attribute("q", Int64, "Integer division quotient", func() {
		Example("default", 4)
	})
	Attribute("r", Int64, "Integer division remainder", func() {
		Example("default", 0)
	})

	Required("q", "r")
})
