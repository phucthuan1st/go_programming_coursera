package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Printf("Enter a string: ")

	fmt.Scan(&str)

	fmt.Printf("Your input string is: %s \n", str)

	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.ContainsRune(str, 'a') {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
