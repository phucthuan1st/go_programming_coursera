package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name    string
	address string
}

func main() {
	var person Person

	fmt.Print("Enter name: ")
	fmt.Scan(&person.name)

	fmt.Print("Enter address: ")
	fmt.Scan(&person.address)

	person_map := make(map[string]string)
	person_map["name"] = person.name
	person_map["address"] = person.address

	data, err := json.Marshal(person_map)

	if err != nil {
		fmt.Printf("Error marshalling person map %s", err)
		fmt.Println()
	} else {
		fmt.Println(string(data))
	}
}
