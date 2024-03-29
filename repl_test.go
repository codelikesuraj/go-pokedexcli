package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "Not YOUR AVERAGE joe",
			expected: []string{
				"not",
				"your",
				"average",
				"joe",
			},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		assert.EqualValues(t, c.expected, actual)
	}
}
