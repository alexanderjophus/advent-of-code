package day5_test

import (
	"testing"

	"github.com/trelore/advent-of-code/go/2018/day5"
)

func TestSample(t *testing.T) {
	actual := day5.Day5([]int{0, 3, 0, 1, -3})
	if actual != 5 {
		t.Error("want 5, got", actual)
	}
}

func TestSampleBonus(t *testing.T) {
	actual := day5.Day5Bonus([]int{0, 3, 0, 1, -3})
	if actual != 10 {
		t.Error("want 10, got", actual)
	}
}
