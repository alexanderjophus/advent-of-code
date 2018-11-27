package day6_test

import (
	"testing"

	"github.com/alexanderjosephtrelore/advent-of-code/day6"
)

func TestDay6(t *testing.T) {
	td := []struct {
		name     string
		f        func([]int) int
		input    []int
		expected int
	}{
		{"0270", day6.Day6, []int{0, 2, 7, 0}, 5},
		{"2412", day6.Day6, []int{2, 4, 1, 2}, 4},
		{"0270", day6.Day6Bonus, []int{0, 2, 7, 0}, 4},
		{"2412", day6.Day6Bonus, []int{2, 4, 1, 2}, 4},
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

func TestDay6Cycle(t *testing.T) {
	td := []struct {
		name     string
		f        func([]int) []int
		input    []int
		expected []int
	}{
		{"0270", day6.Day6Cycle, []int{0, 2, 7, 0}, []int{2, 4, 1, 2}},
		{"2412", day6.Day6Cycle, []int{2, 4, 1, 2}, []int{3, 1, 2, 3}},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			r := tc.f(tc.input)
			if len(tc.expected) != len(r) {
				t.Errorf("got %d, want %d", r, tc.expected)
			}
			for i := range tc.expected {
				if tc.expected[i] != r[i] {
					t.Errorf("element %d, got %d, want %d", i, r[i], tc.expected[i])
				}
			}
		})
	}
}

func BenchmarkDay6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day6.Day6([]int{0, 2, 7, 0})
	}
}
