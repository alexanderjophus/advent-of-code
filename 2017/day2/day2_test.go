package main

import (
	"testing"
)

func TestRangeLineOne(t *testing.T) {
	actual := findRange([]int{5,1,9,5})
	if actual != 8 {
		t.Error("Expected 8, got ", actual)
	}
}

func TestRangeLineTwo(t *testing.T) {
	actual := findRange([]int{7,5,3})
	if actual != 4 {
		t.Error("Expected 4, got ", actual)
	}
}

func TestRangeLineThree(t *testing.T) {
	actual := findRange([]int{2,4,6,8})
	if actual != 6 {
		t.Error("Expected 6, got ", actual)
	}
}

func TestRangeFile1(t *testing.T) {
	actualSum, actualDivisor := dayTwo("test1.txt")
	if actualSum != 18 {
		t.Error("Expected 18, got ", actualSum)
	}
	if actualDivisor != 9{
		t.Error("Expected 9, got ", actualDivisor)
	}
}

func TestDivisibleLineOne(t *testing.T) {
	actual := findAlt([]int{5, 9, 2, 8})
	if actual != 4 {
		t.Error("Expected 4, got ", actual)
	}
}

func TestDivisibleLineTwo(t *testing.T) {
	actual := findAlt([]int{9, 4, 7, 3})
	if actual != 3 {
		t.Error("Expected 3, got ", actual)
	}
}

func TestDivisibleLineThree(t *testing.T) {
	actual := findAlt([]int{3, 8, 6, 5})
	if actual != 2 {
		t.Error("Expected 2, got ", actual)
	}
}
