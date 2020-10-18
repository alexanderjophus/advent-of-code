package day1_test

import (
	"strings"
	"testing"

	"github.com/trelore/advent-of-code/go/2018/day1"
)

func TestDayOne(t *testing.T) {
	td := []struct {
		name     string
		f        func([]string) int
		expected int
	}{
		{"1122", day1.DayOne, 3},
		{"1111", day1.DayOne, 4},
		{"1234", day1.DayOne, 0},
		{"91212129", day1.DayOne, 9},
		{"1212", day1.DayOneBonus, 6},
		{"1221", day1.DayOneBonus, 0},
		{"123425", day1.DayOneBonus, 4},
		{"123123", day1.DayOneBonus, 12},
		{"12131415", day1.DayOneBonus, 4},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.f(strings.Split(tc.name, ""))
			if tc.expected != res {
				t.Errorf("got %d, want %d", res, tc.expected)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day1.DayOne(strings.Split("91212129", ""))
	}
}

func BenchmarkBonus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day1.DayOneBonus(strings.Split("12131415", ""))
	}
}
