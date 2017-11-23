package main

import (
	"fmt"

	"./strategy"
)

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

func main() {
	add := strategy.Operation{Addition{}}
	fmt.Println("3 + 5 =", add.Operate(3, 5))

	mult := strategy.Operation{Multiplication{}}

	fmt.Println("3 * 5 =", mult.Operate(3, 5))
}
