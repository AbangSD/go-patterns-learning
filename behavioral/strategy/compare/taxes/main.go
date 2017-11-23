package main

import (
	"fmt"

	"./taxes"
)

type Chinese struct {
	A float32
	B float32
	C float32
	D float32
}

func (c Chinese) Taxes() float32 {
	return 0.001*c.A + 0.002*c.B + 0.003*c.C + 0.0004*c.D
}

type American struct {
	A float32
	B float32
	C float32
}

func (a American) Taxes() float32 {
	return 0.0015 * (a.A + a.B + a.C)
}

func main() {
	c := Chinese{A: 100, B: 100, C: 100}
	ct := taxes.People{c}
	a := American{A: 100, B: 100, C: 100}
	at := taxes.People{a}
	fmt.Printf("Chinese %.2f, American %.2f\n", ct.Taxes(), at.Taxes())
}
