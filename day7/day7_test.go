package day7_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/alexanderjosephtrelore/advent-of-code/day7"
)

func TestDay7(t *testing.T) {
	td := []struct {
		name     string
		input    string
		expected string
	}{
		{"given example", "test.txt", "tknk"},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.input)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			actual := day7.Day7(file)
			if tc.expected != actual {
				t.Errorf("want %s, got %s", tc.expected, actual)
			}
		})
	}
}
