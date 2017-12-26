package main

import (
	"fmt"

	"github.com/AbangSD/go-patterns-learning/behavioral/strategy/compare/coffee/coffee"
)

func NormalCoffee(cs coffee.Coffee) coffee.Coffee {
	return coffee.Coffee{
		Name:  "Coffee",
		Price: 5,
	}
}

func AddMilk(cs coffee.Coffee) coffee.Coffee {
	return coffee.Coffee{
		Name:  cs.Name + " + Milk",
		Price: cs.Price + 2.5,
	}
}

func AddSugar(cs coffee.Coffee) coffee.Coffee {
	return coffee.Coffee{
		Name:  cs.Name + " + Sugar",
		Price: cs.Price + 1.5,
	}
}

func main() {
	var c coffee.Coffee

	mc := coffee.Make(NormalCoffee, AddMilk)
	sc := coffee.Make(NormalCoffee, AddSugar)
	smc := coffee.Make(coffee.Make(NormalCoffee, AddMilk), AddSugar)

	fmt.Printf("milk coffee %v\nsugar milk %v\nsugar milk coffee %v", mc(c), sc(c), smc(c))
}
