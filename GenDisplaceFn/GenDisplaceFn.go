package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration float64, initial_velocity float64, initial_displacement float64) func(time float64) float64 {
	fn := func(time float64) float64 {
		return (0.5*acceleration*math.Pow(time, 2) + initial_velocity*time + initial_displacement)
	}

	return fn
}

func main() {
	var a, v_0, s_0 float64
	fmt.Println("Enter a, v_0, s_0: ")
	num, err := fmt.Scanf("%f %f %f", &a, &v_0, &s_0)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else if num > 3 {
		fmt.Printf("Exceeded number of arguments")
	} else {
		DisplaceFn := GenDisplaceFn(a, v_0, s_0)

		var t float64
		fmt.Print("Enter value of time: ")
		fmt.Scan(&t)

		displace := DisplaceFn(t)

		fmt.Printf("Displace at time %v is %v\n", t, displace)
	}
}
