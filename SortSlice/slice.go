package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 0, 3)

	var temp string

	for {
		fmt.Print("Enter the next value: ")
		fmt.Scan(&temp)

		if strings.EqualFold("X", temp) {
			fmt.Println("Program end with X signal!")
			break
		}

		num, err := strconv.Atoi(temp)

		if err == nil {
			slice = append(slice, num)
			sort.Slice(slice, func(i, j int) bool {
				return slice[i] < slice[j]
			})

			fmt.Print("Current ascending slice: ")
			for _, value := range slice {
				fmt.Printf("%d ", value)
			}

			fmt.Println()
		} else {
			fmt.Printf("Bad integer input %s. Please input an integer or X to quit\n", err)
			continue
		}
	}
}
