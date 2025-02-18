package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input           string
		expectedCommand string
		expectedArgs    []string
	}{
		"white spaces": {
			input:           "  hello  world  ",
			expectedCommand: "hello",
			expectedArgs:    []string{"world"},
		},
		// add more cases here
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			cmd, args := cleanInput(tc.input)
			assert.Equal(t, tc.expectedCommand, cmd)
			assert.Equal(t, tc.expectedArgs, args)
		})
	}
}
