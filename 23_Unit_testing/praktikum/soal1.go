package calculator

import "errors"

// Addition function returns the result of a + b.
func Addition(a, b float64) float64 {
	return a + b
}

// Subtraction function returns the result of a - b.
func Subtraction(a, b float64) float64 {
	return a - b
}

// Division function returns the result of a / b, unless b is zero in which case it returns an error.
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Multiplication function returns the result of a * b.
func Multiplication(a, b float64) float64 {
	return a * b
}
