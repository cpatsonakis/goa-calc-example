package calcapi_test

import (
	"context"
	"log"

	calcapi "github.com/cpatsonakis/goa-calc-example/goa-calc"
	"github.com/cpatsonakis/goa-calc-example/goa-calc/errorformat"
	calc "github.com/cpatsonakis/goa-calc-example/goa-calc/gen/calc"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"goa.design/goa/v3/middleware"
)

var _ = Describe("CalcEndpoint", Ordered, func() {
	var (
		parentCtx          context.Context
		calcendpoints      *calc.Endpoints
		calcEndpointClient *calc.Client
	)

	BeforeAll(func() {
		parentCtx = context.TODO()
		calcsrvc := calcapi.NewCalc(log.Default())
		calcendpoints = calc.NewEndpoints(calcsrvc)
		calcEndpointClient = calc.NewClient(calcendpoints.Multiply)
	})

	Describe("Execute endpoint multiplications", func() {

		It("should compute correct result for all table inputs", func() {
			for _, mulTableEntry := range mulTestTable {
				ctx := context.WithValue(parentCtx, middleware.RequestIDKey, uuid.NewString())
				output, err := calcEndpointClient.Multiply(ctx, &mulTableEntry.input)
				Expect(err).To(BeNil())
				Expect(output).To(Equal(mulTableEntry.expectedOutput))
			}
		})

		It("should fail because parameter is missing from context", func() {
			dummyCtx := context.TODO()
			output, err := calcEndpointClient.Multiply(dummyCtx, &mulTestTable[0].input)
			Expect(output).To(BeEmpty())
			Expect(err).To(Not(BeNil()))
			Expect(err).To(BeAssignableToTypeOf(&calc.ErrorResultType{}))
			errorResult := err.(*calc.ErrorResultType)
			Expect(errorResult.Name).To(Equal(errorformat.InternalServerErrorName))
		})
	})
})
