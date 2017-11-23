package main

import (
	"fmt"

	"./car"
)

type Car struct {
	color  car.Color
	speed  car.Speed
	wheels car.Wheels
}

func (c Car) Paint(color car.Color) Car {
	c.color = color
	return c
}

func (c Car) TopSpeed(speed car.Speed) Car {
	c.speed = speed
	return c
}

func (c Car) Wheels(wheels car.Wheels) Car {
	c.wheels = wheels
	return c
}

func (c Car) Build() car.Interface {
	return c
}

func (c Car) Drive() error {
	fmt.Println(c.color, c.speed, c.wheels, "car drive")
	return nil
}

func (c Car) Stop() error {
	fmt.Println(c.color, c.speed, c.wheels, "car stop")
	return nil
}

func main() {
	assembly := Car{}.Paint(car.RedColor)

	familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	sportsCar.Stop()
}
