package day3_test

import (
	"strconv"
	"testing"

	"github.com/alexanderjosephtrelore/advent-of-code/day3"
)

func TestDay3(t *testing.T) {
	td := []struct {
		name     int
		f        func(float64) int
		expected int
	}{
		{1, day3.Day3, 0},
		{12, day3.Day3, 3},
		{23, day3.Day3, 2},
		{1024, day3.Day3, 31},
	}

	for _, tc := range td {
		t.Run(strconv.Itoa(tc.name), func(t *testing.T) {
			res := tc.f(float64(tc.name))
			if res != tc.expected {
				t.Errorf("got %d, want %d", res, tc.expected)
			}
		})
	}
}
