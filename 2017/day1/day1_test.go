package main

import (
	"testing"
	"strings"
)

func TestOne(t *testing.T) {
	actual := dayOne(strings.Split("1122", ""))
	if actual != 3 {
		t.Error("Expected 3, got ", actual)
	}
}

func TestTwo(t *testing.T) {
	actual := dayOne(strings.Split("1111", ""))
	if actual != 4 {
		t.Error("Expected 4, got ", actual)
	}
}

func TestThree(t *testing.T) {
	actual := dayOne(strings.Split("1234", ""))
	if actual != 0 {
		t.Error("Expected 0, got ", actual)
	}
}

func TestFour(t *testing.T) {
	actual := dayOne(strings.Split("91212129", ""))
	if actual != 9 {
		t.Error("Expected 9, got ", actual)
	}
}

func BenchmarkOne(b *testing.B) {
	for i := 0 ; i <b.N ; i++ {
		dayOne(strings.Split("1122", ""))
	}
}
