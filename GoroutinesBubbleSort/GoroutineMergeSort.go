package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func mergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	var left, right []int

	go func() {
		defer wg.Done()
		left = mergeSort(slice[:mid]) // Recursive call for sorting the left half
	}()

	go func() {
		defer wg.Done()
		right = mergeSort(slice[mid:]) // Recursive call for sorting the right half
	}()

	wg.Wait()

	return merge(left, right) // Merge the sorted left and right halves
}

func mergeSortConcurrent(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	partitionSize := (len(slice) + 3) / 4 // Calculate partition size rounding up
	partitions := make([][]int, 4)

	var wg sync.WaitGroup
	wg.Add(4)

	for i := 0; i < 4; i++ {
		go func(index int) {
			defer wg.Done()

			start := index * partitionSize
			end := start + partitionSize

			if end > len(slice) {
				end = len(slice)
			}

			partition := slice[start:end]
			fmt.Printf("Sorting partition %d: %v\n", index+1, partition) // Print the partition being sorted
			partitions[index] = mergeSort(partition)                     // Sort the partition

			fmt.Printf("Sorted result of partition %d: %v\n", index+1, partitions[index]) // Print the sorted result of the partition
		}(i)
	}

	wg.Wait()

	// Merge the sorted partitions
	result := merge(partitions[0], partitions[1])
	result = merge(result, partitions[2])
	result = merge(result, partitions[3])

	return result
}

func main() {
	fmt.Println("Enter a series of integers (0 < x < 100) separated by spaces. Press Enter when finished:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	sliceStr := strings.Split(input, " ")
	slice := make([]int, 0, len(sliceStr))

	for _, str := range sliceStr {
		num, _ := strconv.Atoi(str)
		slice = append(slice, num)
	}

	fmt.Println("Processing...")
	fmt.Println()

	sortedSlice := mergeSortConcurrent(slice)

	fmt.Printf("\nSorted slice: %v\n", sortedSlice)
}
