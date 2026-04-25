package main

import (
	"testing"
)

func TestRepl(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "	Hello	World		",
			expected: []string{"hello", "world"},
		},
		{
			input:    "POKEDEX UPPERCASE",
			expected: []string{"pokedex", "uppercase"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v, got: %v", expectedWord, word)
			}
		}
	}
}
