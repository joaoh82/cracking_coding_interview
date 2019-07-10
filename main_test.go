package main

import "testing"

func TestStringHasUniqueChars(t *testing.T) {
	tests := []struct {
		input          string
		expectedResult bool
	}{
		{input: "qwertyuiop", expectedResult: true},
		{input: "hello", expectedResult: false},
	}

	for _, test := range tests {
		result := IsStringCharsUnique(test.input)
		if result != test.expectedResult {
			t.Fatalf("Got %v, Expected %v", result, test.expectedResult)
		}
	}
}
