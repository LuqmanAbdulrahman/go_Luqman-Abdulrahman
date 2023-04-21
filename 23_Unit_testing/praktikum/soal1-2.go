package main

import (
	"testing"
)

type Test struct {
	a, b, expected int
}

func TestAddition(t *testing.T) {
	tests := []Test{
		{2, 3, 5},
		{0, 0, 0},
		{-5, 2, -3},
	}

	for _, test := range tests {
		result := Addition(test.a, test.b)
		if result != test.expected {
			t.Errorf("Addition(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestSubtraction(t *testing.T) {
	tests := []Test{
		{5, 3, 2},
		{0, 0, 0},
		{2, 5, -3},
	}

	for _, test := range tests {
		result := Subtraction(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtraction(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMultiplication(t *testing.T) {
	tests := []Test{
		{2, 3, 6},
		{0, 0, 0},
		{-5, 2, -10},
	}

	for _, test := range tests {
		result := Multiplication(test.a, test.b)
		if result != test.expected {
			t.Errorf("Multiplication(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestDivision(t *testing.T) {
	tests := []Test{
		{6, 3, 2},
		{0, 5, 0},
		{-10, 2, -5},
	}

	for _, test := range tests {
		result, err := Division(test.a, test.b)
		if err != nil {
			t.Errorf("Division(%d, %d) returned an error: %v", test.a, test.b, err)
		} else if result != test.expected {
			t.Errorf("Division(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestDivisionByZero(t *testing.T) {
	_, err := Division(10, 0)
	if err == nil {
		t.Errorf("Division(10, 0) did not return an error")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		// Generate code coverage report
		coverage := testing.Coverage()
		if coverage < 1 {
			panic("Code coverage is not 100%")
		}
	}
}
