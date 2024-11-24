package main

import "testing"

// addTest represents a test case for the Add function
type addTest struct {
	arg1, arg2, expected int
}

// Define test cases for the Add function
var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		output := Add(test.arg1, test.arg2)
		if output != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.arg1, test.arg2, output, test.expected)
		}
	}
}
