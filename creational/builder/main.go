package main

import (
	"fmt"

	"./builder"
)

type Car struct {
	brand builder.Brand
	color builder.Color
	size  builder.Size
}

func (c *Car) Brand(brand builder.Brand) *Car {
	c.brand = brand
	return c
}

func (c *Car) Color(color builder.Color) *Car {
	c.color = color
	return c
}

func (c *Car) Size(size builder.Size) *Car {
	c.size = size
	return c
}

func (c *Car) Build() builder.Product {
	return c
}

func (c *Car) Drive() error {
	fmt.Printf("A %s %s %s car is driving.\n", c.size, c.color, c.brand)
	return nil
}

func (c *Car) Stop() error {
	fmt.Printf("A %s %s %s car is stopped.\n", c.size, c.color, c.brand)
	return nil
}

func main() {
	var c = new(Car)

	c.Brand(builder.BMW).Color(builder.Green).Size(builder.Big)
	bmw := c.Build()
	bmw.Drive()
	bmw.Stop()

	c.Brand(builder.MB).Color(builder.Red).Size(builder.Small)
	mb := c.Build()
	mb.Drive()
	mb.Stop()
}
