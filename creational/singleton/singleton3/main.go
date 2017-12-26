package main

import (
	"fmt"

	"github.com/AbangSD/go-patterns-learning/creational/singleton/singleton3/singleton"
)

func main() {
	s1 := singleton.New()

	s1["this"] = "s1"

	s2 := singleton.New()

	fmt.Println("s2[\"this\"] is", s2["this"])

	s2["this"] = "s2"

	fmt.Println("s1[\"this\"] is", s1["this"])
}
