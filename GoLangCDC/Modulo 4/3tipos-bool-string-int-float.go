package main

import (
	"fmt"
)

func main() {
	// bool (true/false)
	fmt.Printf("Type: %T - Value: %v\n", true, true)
	fmt.Printf("Type: %T - Value: %v\n", false, false)

	// string - sequência de bytes
	fmt.Printf("Type: %T - Value: %v\n", "bruno12", "bruno12")
	fmt.Printf("Type: %T - Value: %v\n", "abcabc", "abcabc")

	// int
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)
	fmt.Printf("Type: %T - Value: %v\n", 123, 123)

	// float (float64/float32) - decimal
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)
	fmt.Printf("Type: %T - Value: %v\n", 3.14, 3.14)
}
