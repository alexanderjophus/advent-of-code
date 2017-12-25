package main

import (
	"testing"
	"strings"
)

func TestLineOne(t *testing.T) {
	passphrase := "aa bb cc dd ee"
	actual := isValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestLineTwo(t *testing.T) {
	passphrase := "aa bb cc dd aa"
	actual := isValidPassphrase(strings.Fields(passphrase))
	if actual != false {
		t.Error("Expected passphrase to be not valid:", passphrase)
	}
}

func TestLineThree(t *testing.T) {
	passphrase := "aa bb cc dd aaa"
	actual := isValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestFile(t *testing.T) {
	file := "test.txt"
	actual := countValidPassphrases(file, isValidPassphrase)
	if actual != 2 {
		t.Error("Expected 2 valid passphrases got: ", actual)
	}
}

func BenchmarkValidity(b *testing.B) {
	passphrase := "aa bb cc dd ee"
	for i := 0 ; i < b.N ; i++ {
		isValidPassphrase(strings.Fields(passphrase))
	}
}

func TestPartTwoLineOne(t *testing.T) {
	passphrase := "abcde fghij"
	actual := isOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestPartTwoLineTwo(t *testing.T) {
	passphrase := "abcde xyz ecdab"
	actual := isOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != false {
		t.Error("Expected passphrase to be not valid:", passphrase)
	}
}

func TestPartTwoLineThree(t *testing.T) {
	passphrase := "a ab abc abd abf abj"
	actual := isOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestPartTwoLineFour(t *testing.T) {
	passphrase := "iiii oiii ooii oooi oooo"
	actual := isOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestPartTwoLineFive(t *testing.T) {
	passphrase := "oiii ioii iioi iiio"
	actual := isOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != false {
		t.Error("Expected passphrase to be not valid:", passphrase)
	}
}
