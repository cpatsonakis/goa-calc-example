package calcapi_test

import (
	"context"
	"log"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"goa.design/goa/v3/middleware"

	calcapi "github.com/cpatsonakis/goa-calc-example/goa-calc"
	"github.com/cpatsonakis/goa-calc-example/goa-calc/errorformat"
	"github.com/cpatsonakis/goa-calc-example/goa-calc/gen/calc"
)

var _ = Describe("Calc Service - Business logic", Ordered, func() {
	var (
		parentCtx context.Context
		calcsrvc  calc.Service
	)

	BeforeAll(func() {
		parentCtx = context.TODO()
		calcsrvc = calcapi.NewCalc(log.Default())
	})

	Describe("Execute service multiplications", func() {
		It("should compute correct result for all table inputs", func() {
			for _, mulTableEntry := range mulTestTable {
				ctx := context.WithValue(parentCtx, middleware.RequestIDKey, uuid.NewString())
				output, err := calcsrvc.Multiply(ctx, &mulTableEntry.input)
				Expect(err).To(BeNil())
				Expect(output).To(Equal(mulTableEntry.expectedOutput))
			}
		})

		It("should fail because parameter is missing from context", func() {
			dummyCtx := context.TODO()
			output, err := calcsrvc.Multiply(dummyCtx, &mulTestTable[0].input)
			Expect(output).To(BeEmpty())
			Expect(err).To(Not(BeNil()))
			Expect(err).To(BeAssignableToTypeOf(&calc.ErrorResultType{}))
			errorResult := err.(*calc.ErrorResultType)
			Expect(errorResult.Name).To(Equal(errorformat.InternalServerErrorName))
		})
	})
})
