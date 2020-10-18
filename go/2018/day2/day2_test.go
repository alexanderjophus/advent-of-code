package day2_test

import (
	"os"
	"testing"

	"github.com/trelore/advent-of-code/go/2018/day2"
)

func TestDayOne(t *testing.T) {
	td := []struct {
		name     string
		f        func([]int) int
		input    []int
		expected int
	}{
		{"5195", day2.FindRange, []int{5, 1, 9, 5}, 8},
		{"753", day2.FindRange, []int{7, 5, 3}, 4},
		{"2468", day2.FindRange, []int{2, 4, 6, 8}, 6},
		{"5928", day2.FindAlt, []int{5, 9, 2, 8}, 4},
		{"9473", day2.FindAlt, []int{9, 4, 7, 3}, 3},
		{"3865", day2.FindAlt, []int{3, 8, 6, 5}, 2},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.f(tc.input)
			if tc.expected != res {
				t.Errorf("got %d, want %d", res, tc.expected)
			}
		})
	}
}

func TestRangeFile1(t *testing.T) {
	file, _ := os.Open("test1.txt")
	defer file.Close()
	actualSum, actualDivisor := day2.DayTwo(file)
	if actualSum != 18 {
		t.Error("Expected 18, got ", actualSum)
	}
	if actualDivisor != 9 {
		t.Error("Expected 9, got ", actualDivisor)
	}
}
