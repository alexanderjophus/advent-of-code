package solution

import (
	"testing"
	"strings"
)

func TestOne(t *testing.T) {
	actual := DayOne(strings.Split("1122", ""))
	if actual != 3 {
		t.Error("Expected 3, got ", actual)
	}
}

func TestTwo(t *testing.T) {
	actual := DayOne(strings.Split("1111", ""))
	if actual != 4 {
		t.Error("Expected 4, got ", actual)
	}
}

func TestThree(t *testing.T) {
	actual := DayOne(strings.Split("1234", ""))
	if actual != 0 {
		t.Error("Expected 0, got ", actual)
	}
}

func TestFour(t *testing.T) {
	actual := DayOne(strings.Split("91212129", ""))
	if actual != 9 {
		t.Error("Expected 9, got ", actual)
	}
}

func TestOneBonus(t *testing.T) {
	actual := DayOneBonus(strings.Split("1212", ""))
	if actual != 6 {
		t.Error("Expected 6, got ", actual)
	}
}

func TestTwoBonus(t *testing.T) {
	actual := DayOneBonus(strings.Split("1221", ""))
	if actual != 0 {
		t.Error("Expected 0, got ", actual)
	}
}

func TestThreeBonus(t *testing.T) {
	actual := DayOneBonus(strings.Split("123425", ""))
	if actual != 4 {
		t.Error("Expected 4, got ", actual)
	}
}

func TestFourBonus(t *testing.T) {
	actual := DayOneBonus(strings.Split("123123", ""))
	if actual != 12 {
		t.Error("Expected 12, got ", actual)
	}
}

func TestFiveBonus(t *testing.T) {
	actual := DayOneBonus(strings.Split("12131415", ""))
	if actual != 4 {
		t.Error("Expected 4, got ", actual)
	}
}

func Benchmark(b *testing.B) {
	for i := 0 ; i <b.N ; i++ {
		DayOne(strings.Split("91212129", ""))
	}
}

func BenchmarkBonus(b *testing.B) {
	for i := 0 ; i <b.N ; i++ {
		DayOneBonus(strings.Split("12131415", ""))
	}
}
