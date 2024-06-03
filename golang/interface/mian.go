package main

type Operator interface {
	Operate(int, int) int
}

type Add struct{}
type Sub struct{}
type Mul struct{}
type Div struct{}

func (Add) Operate(a, b int) int {
	return a + b
}

func (Sub) Operate(a, b int) int {
	return a - b
}

func (Mul) Operate(a, b int) int {
	return a * b
}

func (Div) Operate(a, b int) int {
	return a / b
}

// op is a interface type
// can be any type that implements the Operator interface
func calculate(op Operator, a, b int) int {
	return op.Operate(a, b)
}

func main() {
	add := Add{}
	sub := Sub{}
	mul := Mul{}
	div := Div{}

	println(calculate(add, 1, 2))
	println(calculate(sub, 1, 2))
	println(calculate(mul, 1, 2))
	println(calculate(div, 1, 2))
}
