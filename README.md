# goa-calc-example
First time learning experience using the Goa framework for designing and developing a simple calculator microservice in Go.


To generate:

```bash
goa gen github.com/cpatsonakis/goa-calc-example/design/goa-calc -o goa-calc
```

To generate the example:

```bash
goa example github.com/cpatsonakis/goa-calc-example/design/goa-calc -o goa-calc
```

The following code was added to the multiplication service implementation:

```go
// Multiply two integers a and b and get the result in the response's body.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplicationPayload) (res string, err error) {
	s.logger.Print("calc.multiply")
	a := big.NewInt(p.A)
	b := big.NewInt(p.B)
	return a.Mul(a, b).String(), nil
}
```

Build the executable of the service:

```bash
go build ./goa-calc/cmd/calc
```

Run the service:

```bash
./calc -debug -http-port 8080
```


Clean-up:

```bash
rm -rf goa-calc/gen
```