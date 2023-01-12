package lib

import (
	"fmt"
)

// Returns the sum of two numbers
func Add(a int, b int) int {
	fmt.Println("add data")
	return a + b
}

// Returns the difference between two numbers
func Subtract(a int, b int) int {
	fmt.Println("substract data")
	return a - b
}
