package main

import (
	"fmt"
)

type Number interface {
	float64 | int64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// Call the SumNumbers function with the ints map
	fmt.Println(SumNumbers(ints))
	fmt.Println(SumNumbers(floats))
}
