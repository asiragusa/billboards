package main

import (
	"strings"
	"testing"

	"bufio"

	"github.com/stretchr/testify/assert"
)

func TestReadInputError(t *testing.T) {
	inputs := []string{
		"",
		"notANumber",
		"0",
		"-1",
		"1\n",
		"2\n1 1 text",
		"20\n" + strings.Repeat("1 1 text\n", 21),
		"21\n" + strings.Repeat("1 1 text\n", 21),
		"1\n1 text",
	}

	for _, input := range inputs {
		scanner := bufio.NewScanner(strings.NewReader(input))
		billboards, err := readInput(scanner)
		if assert.Error(t, err, `"`+input+`"`) {
			assert.Nil(t, billboards)
		}
	}
}

func TestReadInputOK(t *testing.T) {
	text := "1 1 text\n"

	type test struct {
		length int
		input  string
	}
	tests := []test{
		{1, "1\n" + text},
		{1, "1\n1 1 text"},
		{20, "20\n" + strings.Repeat(text, 20)},
		{20, "20\n" + strings.Repeat(text, 19) + "1 1 text"},
	}

	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.input))
		billboards, err := readInput(scanner)
		if assert.NoError(t, err, `"`+test.input+`"`) {
			assert.Len(t, billboards, test.length)
		}
	}
}
