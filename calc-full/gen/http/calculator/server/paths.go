// Code generated by goa v3.10.1, DO NOT EDIT.
//
// HTTP request path constructors for the calculator service.
//
// Command:
// $ goa gen github.com/cpatsonakis/goa-calc-example/design/calc-full -o
// calc-full

package server

import (
	"fmt"
)

// AddCalculatorPath returns the URL path to the calculator service add HTTP endpoint.
func AddCalculatorPath(a int64, b int64) string {
	return fmt.Sprintf("/calculator/add/%v/%v", a, b)
}

// SubtractCalculatorPath returns the URL path to the calculator service subtract HTTP endpoint.
func SubtractCalculatorPath(a int64, b int64) string {
	return fmt.Sprintf("/calculator/sub/%v/%v", a, b)
}

// MultiplyCalculatorPath returns the URL path to the calculator service multiply HTTP endpoint.
func MultiplyCalculatorPath(a int64, b int64) string {
	return fmt.Sprintf("/calculator/mul/%v/%v", a, b)
}

// DivideCalculatorPath returns the URL path to the calculator service divide HTTP endpoint.
func DivideCalculatorPath(a int64, b int64) string {
	return fmt.Sprintf("/calculator/div/%v/%v", a, b)
}
