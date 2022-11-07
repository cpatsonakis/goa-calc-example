package design

import (
	. "goa.design/goa/v3/dsl"
)

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
