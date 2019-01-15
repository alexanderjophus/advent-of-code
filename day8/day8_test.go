package day8_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/trelore/advent-of-code/day8"
)

func TestDay8(t *testing.T) {
	t.Skip("not implemented yet")
	td := []struct {
		name     string
		input    string
		expected int
		err      error
	}{
		{"given example", "test.txt", 1, nil},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.input)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			actual, err := day8.Day8(file)
			if tc.err != err {
				t.Errorf("want %s, got %s", tc.err, err)
			}
			if tc.expected != actual {
				t.Errorf("want %d, got %d", tc.expected, actual)
			}
		})
	}
}
