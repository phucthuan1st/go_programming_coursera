package main

import (
	"fmt"
	"strconv"
)

func swap(slice []int, index int) {
	temp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = temp
}

func BubbleSort(slice []int) {
	fmt.Println("Running bubble sort")
	n := len(slice)

	fmt.Printf(" Current slice: %v\n", slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}

		fmt.Printf("Modified slice: %v\n", slice)
	}

	fmt.Println("Bubble sort completed")
}

func getSliceElement(slice []int) []int {

	fmt.Println("Enter ten integers: ")

	for i := 0; i < 10; i++ {
		var temp string

		_, err := fmt.Scan(&temp)

		if err != nil {
			fmt.Printf("Bad input %s with error %v\n", temp, err)
			continue
		}

		num, err := strconv.Atoi(temp)
		if err != nil {
			fmt.Printf("Bad integer input %s with error %v\n", temp, err)
			continue
		}

		slice = append(slice, num)
	}

	return slice
}

func main() {
	slice := make([]int, 0, 1)

	slice = getSliceElement(slice)

	BubbleSort(slice)

	fmt.Printf("Sorted slice: %v\n", slice)
}
