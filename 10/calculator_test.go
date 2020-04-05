package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 2},
		{a: 2, b: 2, want: 4},
		{a: 6, b: 7, want: 13},
		{a: 0, b: 7, want: 7},
		{a: 6, b: 0, want: 6},
		{a: -1, b: -1, want: -2},
		{a: 3, b: -1, want: 2},
		{a: -3, b: 1, want: -2},
		{a: -3, b: 4, want: 1},
		{a: 3, b: -7, want: -4},
	}
	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 0},
		{a: 2, b: 2, want: 0},
		{a: 6, b: 7, want: -1},
		{a: 0, b: 7, want: -7},
		{a: 6, b: 0, want: 6},
		{a: -1, b: -1, want: 0},
		{a: 3, b: -1, want: 4},
		{a: -3, b: 1, want: -4},
		{a: -3, b: 4, want: -7},
		{a: 3, b: -7, want: 10},
	}
	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Subtract(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 1},
		{a: 2, b: 2, want: 4},
		{a: 6, b: 7, want: 42},
		{a: 0, b: 7, want: 0},
		{a: 6, b: 0, want: 0},
		{a: -1, b: -1, want: 1},
		{a: 3, b: -1, want: -3},
		{a: -3, b: 1, want: -3},
		{a: -3, b: 4, want: -12},
		{a: 3, b: -7, want: -21},
	}
	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Multiply(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b        int
		want        int
		errExpected bool
	}{
		{a: 1, b: 1, want: 1, errExpected: false},
		{a: 2, b: 2, want: 1, errExpected: false},
		{a: 21, b: 7, want: 3, errExpected: false},
		{a: 0, b: 7, want: 0, errExpected: false},
		{a: 12, b: -2, want: -6, errExpected: false},
		{a: -1, b: -1, want: 1, errExpected: false},
		{a: 3, b: -1, want: -3, errExpected: false},
		{a: -3, b: 1, want: -3, errExpected: false},
		{a: -12, b: 6, want: -2, errExpected: false},
		{a: 28, b: -7, want: -4, errExpected: false},
		{a: 1, b: 0, want: 0, errExpected: true},
	}
	for _, testCase := range testCases {
		got, err := calculator.Divide(testCase.a, testCase.b)

		if err == nil && testCase.errExpected {
			t.Fatalf("Expected error for Divide(%d, %d)", testCase.a, testCase.b)
		}

		if testCase.want != got {
			t.Errorf("Divide(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}
