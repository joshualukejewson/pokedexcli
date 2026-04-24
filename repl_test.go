package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " this is a test     ",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "                BULBAsaur           ",
			expected: []string{"bulbasaur"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		fmt.Println(actual, c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths of input and expected do not match.")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("the strings do not match.")
			}
		}
	}
}
