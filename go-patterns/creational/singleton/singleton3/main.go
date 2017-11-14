package main

import (
	"fmt"

	"./singleton"
)

func main() {
	s1 := singleton.New()

	s1["this"] = "s1"

	s2 := singleton.New()

	fmt.Println("s2[\"this\"] is", s2["this"])

	s2["this"] = "s2"

	fmt.Println("s1[\"this\"] is", s1["this"])

	a1 := &s1
	a2 := &s2
	fmt.Println("s1 address is", &a1, "\ns2 address is", &a2)
}

// s2["this"] is s1
// s1["this"] is s2
// s1 address is 0xc42007e030
// s2 address is 0xc42007e038
