package main

import "fmt"

type CoffeeComponent interface {
	Cost() int
	Description() string
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() int {
	return 100
}

func (c *SimpleCoffee) Description() string {
	return "Simple coffee"
}

type MilkDecorator struct {
	coffee CoffeeComponent
}

func NewMilkDecorator(coffee CoffeeComponent) *MilkDecorator {
	return &MilkDecorator{coffee: coffee}
}

func (d *MilkDecorator) Cost() int {
	return d.coffee.Cost() + 50
}

func (d *MilkDecorator) Description() string {
	return d.coffee.Description() + ", with milk"
}

func main() {
	simpleCoffee := &SimpleCoffee{}

	coffeeWithMilk := NewMilkDecorator(simpleCoffee)

	fmt.Printf("Description: %s\n", coffeeWithMilk.Description())
	fmt.Printf("Cost: %d dollars\n", coffeeWithMilk.Cost())
}
