package day4_test

import (
	"os"
	"strings"
	"testing"

	"github.com/alexanderjosephtrelore/advent-of-code/day4"
)

func TestDayOne(t *testing.T) {
	td := []struct {
		name     string
		f        func([]string) bool
		expected bool
	}{
		{"aa bb cc dd ee", day4.IsValidPassphrase, true},
		{"aa bb cc dd aa", day4.IsValidPassphrase, false},
		{"aa bb cc dd aaa", day4.IsValidPassphrase, true},
		{"abcde fghij", day4.IsOrderedValidPassphrase, true},
		{"abcde xyz ecdab", day4.IsOrderedValidPassphrase, false},
		{"a ab abc abd abf abj", day4.IsOrderedValidPassphrase, true},
		{"iiii oiii ooii oooi oooo", day4.IsOrderedValidPassphrase, true},
		{"oiii ioii iioi iiio", day4.IsOrderedValidPassphrase, false},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.f(strings.Fields(tc.name))
			if tc.expected != res {
				t.Errorf("got %v, want %v", res, tc.expected)
			}
		})
	}
}

func TestFile(t *testing.T) {
	file, _ := os.Open("test.txt")
	defer file.Close()
	actual := day4.CountValidPassphrases(file, day4.IsValidPassphrase)
	if actual != 2 {
		t.Error("Expected 2 valid passphrases got: ", actual)
	}
}

func BenchmarkValidity(b *testing.B) {
	passphrase := "aa bb cc dd ee"
	for i := 0; i < b.N; i++ {
		day4.IsValidPassphrase(strings.Fields(passphrase))
	}
}
