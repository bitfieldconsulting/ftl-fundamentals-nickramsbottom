package calculator

import "errors"

// Add the argument integers
func Add(a, b int) int {
	return a + b
}

// Subtract two numbers
func Subtract(a, b int) int {
	return a - b
}

// Multiply two numbers
func Multiply(a, b int) int {
	return a * b
}

// Divide two numbers
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero: %d, %d", a, b)
	}
	return a / b, nil
}
