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
