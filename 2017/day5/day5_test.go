package main

import "testing"

func TestSample(t *testing.T) {
	actual := day5([]int{0, 3, 0, 1, -3})
	if actual != 5 {
		t.Error("Expected 5, got", actual)
	}
}

func TestFromFile(t *testing.T) {
	actual := day5(readFile("test.txt"))
	if actual != 5 {
		t.Error("Expected 5, got,", actual)
	}
}

func TestSampleBonus(t *testing.T) {
	actual := day5Bonus([]int{0, 3, 0, 1, -3})
	if actual != 10 {
		t.Error("Expected 10, got", actual)
	}
}
