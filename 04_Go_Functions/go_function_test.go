package main

import "testing"

func TestMyFunction(t *testing.T) {
	expected := "Elliot Forbes"
	result := MyFunction("Elliot", "Forbes")

	if expected != result {
		t.Error("The result does not match what we expected")
	}
}

func TestMyFunctionWithError(t *testing.T) {
	expected := "Elliot Forbes"
	result, err := MyFunctionWithError("Elliot", "Forbes")
	if err != nil {
		t.Error("Our function Threw an Error")
	}
	if expected != result {
		t.Error("The result does not match what we expected")
	}
}

func TestNegativeMyFunctionWithError(t *testing.T) {
	result, err := MyFunctionWithError("E", "Forbes")
	if err == nil {
		t.Error("Our Function never threw an error")
	}
	if result != "" {
		t.Error("Expected an empty string returned")
	}
}

func TestChallenge01(t *testing.T) {
	var tests = []struct {
		value1   int
		value2   int
		expected int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 2, 2},
		{-5, -3, -8},
		{99999, 1, 100000},
	}

	for _, test := range tests {
		result := Add(test.value1, test.value2)
		if result != test.expected {
			t.Error("Expected output did not match the result")
		}
	}
}
