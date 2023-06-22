package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string

	fmt.Print("Enter data file name: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)

	people := make([]Name, 0, 3)

	if err != nil {
		log.Fatalf("Cannot open %s: %v", filename, err)
	} else {
		fmt.Println("File opened successfully")

		scanner := bufio.NewScanner(file)

		counter := 0

		for scanner.Scan() {
			line := scanner.Text()
			token := strings.Split(line, " ")

			people = append(people, Name{token[0], token[1]})
			counter++
		}

		fmt.Printf("Success reading %d people names:\n", counter)
		for index, name := range people {
			fmt.Printf("%d: %v\n", index+1, name)
		}
	}

	file.Close()
}
