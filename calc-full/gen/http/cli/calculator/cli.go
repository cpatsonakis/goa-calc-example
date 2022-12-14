// Code generated by goa v3.10.1, DO NOT EDIT.
//
// calculator HTTP client CLI support package
//
// Command:
// $ goa gen github.com/cpatsonakis/goa-calc-example/design/calc-full -o
// calc-full

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	calculatorc "github.com/cpatsonakis/goa-calc-example/calc-full/gen/http/calculator/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `calculator (add|subtract|multiply|divide)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` calculator add --a 3 --b 5` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		calculatorFlags = flag.NewFlagSet("calculator", flag.ContinueOnError)

		calculatorAddFlags = flag.NewFlagSet("add", flag.ExitOnError)
		calculatorAddAFlag = calculatorAddFlags.String("a", "REQUIRED", "First operand of addition payload")
		calculatorAddBFlag = calculatorAddFlags.String("b", "REQUIRED", "Second operand of addition payload")

		calculatorSubtractFlags = flag.NewFlagSet("subtract", flag.ExitOnError)
		calculatorSubtractAFlag = calculatorSubtractFlags.String("a", "REQUIRED", "First operand of subtraction payload")
		calculatorSubtractBFlag = calculatorSubtractFlags.String("b", "REQUIRED", "Second operand of subtraction payload")

		calculatorMultiplyFlags = flag.NewFlagSet("multiply", flag.ExitOnError)
		calculatorMultiplyAFlag = calculatorMultiplyFlags.String("a", "REQUIRED", "First operand of multiplication payload")
		calculatorMultiplyBFlag = calculatorMultiplyFlags.String("b", "REQUIRED", "Second operand of multiplication payload")

		calculatorDivideFlags = flag.NewFlagSet("divide", flag.ExitOnError)
		calculatorDivideAFlag = calculatorDivideFlags.String("a", "REQUIRED", "First operand (nominator) of division payload")
		calculatorDivideBFlag = calculatorDivideFlags.String("b", "REQUIRED", "Second operand (denominator) of division payload")
	)
	calculatorFlags.Usage = calculatorUsage
	calculatorAddFlags.Usage = calculatorAddUsage
	calculatorSubtractFlags.Usage = calculatorSubtractUsage
	calculatorMultiplyFlags.Usage = calculatorMultiplyUsage
	calculatorDivideFlags.Usage = calculatorDivideUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "calculator":
			svcf = calculatorFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "calculator":
			switch epn {
			case "add":
				epf = calculatorAddFlags

			case "subtract":
				epf = calculatorSubtractFlags

			case "multiply":
				epf = calculatorMultiplyFlags

			case "divide":
				epf = calculatorDivideFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "calculator":
			c := calculatorc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = calculatorc.BuildAddPayload(*calculatorAddAFlag, *calculatorAddBFlag)
			case "subtract":
				endpoint = c.Subtract()
				data, err = calculatorc.BuildSubtractPayload(*calculatorSubtractAFlag, *calculatorSubtractBFlag)
			case "multiply":
				endpoint = c.Multiply()
				data, err = calculatorc.BuildMultiplyPayload(*calculatorMultiplyAFlag, *calculatorMultiplyBFlag)
			case "divide":
				endpoint = c.Divide()
				data, err = calculatorc.BuildDividePayload(*calculatorDivideAFlag, *calculatorDivideBFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// calculatorUsage displays the usage of the calculator command and its
// subcommands.
func calculatorUsage() {
	fmt.Fprintf(os.Stderr, `The calculator service performs legendary mathematical operations on integers.
Usage:
    %[1]s [globalflags] calculator COMMAND [flags]

COMMAND:
    add: Addition of two integers.
    subtract: Subtraction of two numbers.
    multiply: Multiplication of two numbers.
    divide: Division of two numbers.

Additional help:
    %[1]s calculator COMMAND --help
`, os.Args[0])
}
func calculatorAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calculator add -a INT64 -b INT64

Addition of two integers.
    -a INT64: First operand of addition payload
    -b INT64: Second operand of addition payload

Example:
    %[1]s calculator add --a 3 --b 5
`, os.Args[0])
}

func calculatorSubtractUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calculator subtract -a INT64 -b INT64

Subtraction of two numbers.
    -a INT64: First operand of subtraction payload
    -b INT64: Second operand of subtraction payload

Example:
    %[1]s calculator subtract --a 5 --b 3
`, os.Args[0])
}

func calculatorMultiplyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calculator multiply -a INT64 -b INT64

Multiplication of two numbers.
    -a INT64: First operand of multiplication payload
    -b INT64: Second operand of multiplication payload

Example:
    %[1]s calculator multiply --a 3 --b 5
`, os.Args[0])
}

func calculatorDivideUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calculator divide -a INT64 -b INT64

Division of two numbers.
    -a INT64: First operand (nominator) of division payload
    -b INT64: Second operand (denominator) of division payload

Example:
    %[1]s calculator divide --a 8 --b 2
`, os.Args[0])
}
