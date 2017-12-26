package main

import (
	"fmt"
)

type Object func(int) int

func Decorate(fsource Object, fextend Object) Object {
	return func(n int) int {
		result := fextend(fsource(n))

		return result
	}
}

func AddOne(n int) int {
	return n + 1
}

func Double(n int) int {
	return n * 2
}

func main() {
	f1 := Decorate(AddOne, Double)
	fmt.Println(f1(5))
	f2 := Decorate(f1, AddOne)
	fmt.Println(f2(5))
}
