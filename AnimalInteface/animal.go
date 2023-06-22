package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Speak()
	Eat()
	Move()
}

type Cow struct {
	name string
}

func (c Cow) Speak() {
	fmt.Printf("%s speaks moo\n", c.name)
}

func (c Cow) Move() {
	fmt.Printf("%s moves by walk\n", c.name)
}

func (c Cow) Eat() {
	fmt.Printf("%s consumes grass\n", c.name)
}

type Bird struct {
	name string
}

func (b Bird) Speak() {
	fmt.Printf("%s speaks peep\n", b.name)
}

func (b Bird) Move() {
	fmt.Printf("%s moves by fly\n", b.name)
}

func (b Bird) Eat() {
	fmt.Printf("%s consumes worms\n", b.name)
}

type Snake struct {
	name string
}

func (s Snake) Speak() {
	fmt.Printf("%s speaks hsss\n", s.name)
}

func (s Snake) Move() {
	fmt.Printf("%s moves by slither\n", s.name)
}

func (s Snake) Eat() {
	fmt.Printf("%s consumes mice\n", s.name)
}

func main() {

	var command string
	var args [2]string

	animals := make(map[string]Animal)

	fmt.Println("Welcome to Animal Dictionary")

	for {
		fmt.Print("> ")

		num, err := fmt.Scanf("%s %s %s", &command, &args[0], &args[1])

		if err != nil {
			fmt.Printf("Bad command input: %s\n", err)
			continue
		}

		if num != 3 {
			fmt.Printf("Invalid number of arguments: expected %d, got %d\n", 3, num)
			continue
		}

		command = strings.ToLower(command)
		args[0] = strings.ToLower(args[0])
		args[1] = strings.ToLower(args[1])

		if command == "newanimal" {
			switch args[1] {
			case "bird":
				{
					animals[args[0]] = Bird{name: args[0]}
				}
			case "cow":
				{
					animals[args[0]] = Cow{name: args[0]}
				}

			case "snake":
				{
					animals[args[0]] = Snake{name: args[0]}
				}
			default:
				{
					fmt.Printf("Invalid type of animals: expected [bird, cow, snake], got %s\n", args[0])
					continue
				}
			}

			fmt.Printf("The %s %s added to dictionary!!!\n", args[1], args[0])
		} else if command == "query" {
			animal, is_exist := animals[args[0]]

			if !is_exist {
				fmt.Printf("%s is not exist\n", args[0])
				continue
			}

			switch args[1] {
			case "speak":
				{
					animal.Speak()
				}
			case "eat":
				{
					animal.Eat()
				}

			case "move":
				{
					animal.Move()
				}
			default:
				{
					fmt.Printf("Invalid type of infomation: expected [eat, move, speak], got %s\n", args[1])
				}
			}
		} else {
			fmt.Println("Invalid command: ", command, "! Must be newanimal or query")
		}
	}
}
