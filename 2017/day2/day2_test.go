package main

import (
	"testing"
)

func TestSumLineOne(t *testing.T) {
	actual := sumLine([]int{5,1,9,5})
	if actual != 8 {
		t.Error("Expected 8, got ", actual)
	}
}

func TestSumLineTwo(t *testing.T) {
	actual := sumLine([]int{7,5,3})
	if actual != 4 {
		t.Error("Expected 4, got ", actual)
	}
}

func TestSumLineThree(t *testing.T) {
	actual := sumLine([]int{2,4,6,8})
	if actual != 6 {
		t.Error("Expected 6, got ", actual)
	}
}

func TestFile1(t *testing.T) {
	actual := dayTwo("test1.txt")
	if actual != 18 {
		t.Error("Expected 18, got ", actual)
	}
}

