package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (a *Animal) Init(specie string) {
	specie = strings.ToLower(specie)
	a.name = specie

	switch specie {
	case "cow":
		{
			a.food = "grass"
			a.locomotion = "walk"
			a.sound = "moo"
		}
	case "snake":
		{
			a.food = "mice"
			a.locomotion = "slither"
			a.sound = "hssss"
		}
	case "bird":
		{
			a.food = "worms"
			a.locomotion = "fly"
			a.sound = "peep"
		}
	default:
		{
			a.food = "unknown"
			a.locomotion = "unknown"
			a.sound = "unknown"
		}
	}
}

func (a *Animal) Eat() {
	fmt.Printf("%s eats %s\n", a.name, a.food)
}

func (a *Animal) Move() {
	fmt.Printf("%s moves by %s\n", a.name, a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Printf("%s speaks %s\n", a.name, a.sound)
}

func main() {
	var specie string
	var command string

	fmt.Println("Welcome to Animal Dictionary")

	for {
		fmt.Print("> ")

		num, err := fmt.Scanf("%s %s", &specie, &command)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		if num != 2 {
			fmt.Println("Invalid command")
			continue
		}

		var animal Animal
		animal.Init(specie)

		command = strings.ToLower(command)

		switch command {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}
	}
}
