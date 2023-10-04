package main

import (
	"fmt"
)

type IPizza interface {
	getPrice() int32
	getInfo() string
}

type Pizza struct {
}

func (p *Pizza) getPrice() int32 {
	return 2000
}

func (p *Pizza) getInfo() string {
	return "Just Pizza"
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int32 {
	price := c.pizza.getPrice() + 300
	return price
}

func (c *CheeseTopping) getInfo() string {
	return c.pizza.getInfo() + ", Cheese Topping"
}

type PepperoniTopping struct {
	pizza IPizza
}

func (c *PepperoniTopping) getPrice() int32 {
	price := c.pizza.getPrice() + 700
	return price
}

func (c *PepperoniTopping) getInfo() string {
	return c.pizza.getInfo() + ", Pepperoni Topping"
}

func main() {
	pizza := &PepperoniTopping{&CheeseTopping{&Pizza{}}}

	fmt.Println(pizza.getPrice())
	fmt.Println(pizza.getInfo())
}
