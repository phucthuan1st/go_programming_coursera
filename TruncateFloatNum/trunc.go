// package main

// import "fmt"

// func main() {
// 	fmt.Printf("Input a number: ")
// 	var number float64
// 	fmt.Scan(&number)

// 	fmt.Printf("The truncated number is: %d \n", int(number))
// }

package main

import "fmt"

func main() {
	var userNum float32
	var intPart int

	fmt.Printf("Enter a floating point number: ")

	count, ferr := fmt.Scan(&userNum)
	intPart = int(userNum)

	if count > 0 {
		fmt.Printf("Original Float: %f\n", userNum)
		fmt.Printf("Truncated to Integer: %d\n", intPart)
	} else if ferr != nil {
		fmt.Printf("Error: %s\n", ferr)
	} else {
		fmt.Printf("Bad float entered\n")
	}
}
