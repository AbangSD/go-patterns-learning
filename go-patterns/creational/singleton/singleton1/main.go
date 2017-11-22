package main

import (
	"fmt"

	"./singleton"
)

func main() {
	s1 := singleton.New()

	s1.N = 1

	s2 := singleton.New()

	fmt.Println("s1 is", s1, "s2 is", s2)

	s2.N = 2

	fmt.Println("s1 is", s1, "s2 is", s2)
}
