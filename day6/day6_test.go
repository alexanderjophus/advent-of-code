package day6_test

import (
	"testing"

	"github.com/alexanderjosephtrelore/advent-of-code/day6"
)

func TestDay6(t *testing.T) {
	t.Skip()
	td := []struct {
		name     string
		f        func([]int) int
		input    []int
		expected int
	}{
		{"0270", day6.Day6, []int{0, 2, 7, 0}, 5},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			r := tc.f(tc.input)
			if tc.expected != r {
				t.Errorf("got %d, want %d", r, tc.expected)
			}
		})
	}
}
