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

	a1 := &s1
	a2 := &s2
	fmt.Println("s1 address is", &a1, "\ns2 address is", &a2)
}

// s1 is &{1} s2 is &{1}
// s1 is &{2} s2 is &{2}
// s1 address is 0xc42000c040
// s2 address is 0xc42000c048
