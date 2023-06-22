package main

import (
	"fmt"
	"time"
)

var money int

func withdraw() {
	fmt.Printf("Current amount of money before withdraw: %d\n", money)
	money -= 15
	fmt.Printf("Current amount of money after withdraw: %d\n", money)
	fmt.Println()
}

func deposit() {
	fmt.Printf("Current amount of money before deposit: %d\n", money)
	money += 12
	fmt.Printf("Current amount of money after deposit: %d\n", money)
	fmt.Println()
}

func main() {
	money = 100

	go func() {
		for i := 0; i < 5; i++ {
			withdraw()
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			deposit()
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Printf("Current amount of money: %d\n", money)
}
