package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"white spaces": {
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, cleanInput(tc.input))
		})
	}
}
