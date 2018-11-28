package day7_test

import (
	"testing"

	"github.com/alexanderjosephtrelore/advent-of-code/day7"
)

func TestDay7(t *testing.T) {
	td := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{"given example", "test.txt", "tknk", nil},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := day7.Day7(tc.input)
			if tc.err != err {
				t.Errorf("want %s, got %s", tc.err, err)
			}
			if tc.expected != actual {
				t.Errorf("want %s, got %s", tc.expected, actual)
			}
		})
	}
}
