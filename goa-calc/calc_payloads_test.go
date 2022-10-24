package calcapi_test

import calc "github.com/cpatsonakis/goa-calc-example/goa-calc/gen/calc"

type multiplyServiceTest struct {
	input          calc.MultiplicationPayload
	expectedOutput string
}

var mulTestTable = []multiplyServiceTest{
	{
		input: calc.MultiplicationPayload{
			A: 1,
			B: 1,
		},
		expectedOutput: "1",
	},
	{
		input: calc.MultiplicationPayload{
			A: 5,
			B: 2,
		},
		expectedOutput: "10",
	},
	{
		input: calc.MultiplicationPayload{
			A: -1,
			B: 5,
		},
		expectedOutput: "-5",
	},
	{
		input: calc.MultiplicationPayload{
			A: -1,
			B: -51,
		},
		expectedOutput: "51",
	},
}
