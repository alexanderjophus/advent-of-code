package solution

import (
	"testing"
	"strings"
	"os"
)

func TestLineOne(t *testing.T) {
	passphrase := "aa bb cc dd ee"
	actual := IsValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestLineTwo(t *testing.T) {
	passphrase := "aa bb cc dd aa"
	actual := IsValidPassphrase(strings.Fields(passphrase))
	if actual != false {
		t.Error("Expected passphrase to be not valid:", passphrase)
	}
}

func TestLineThree(t *testing.T) {
	passphrase := "aa bb cc dd aaa"
	actual := IsValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestFile(t *testing.T) {
	file, _ := os.Open("test.txt")
	defer file.Close()
	actual := CountValidPassphrases(file, IsValidPassphrase)
	if actual != 2 {
		t.Error("Expected 2 valid passphrases got: ", actual)
	}
}

func BenchmarkValidity(b *testing.B) {
	passphrase := "aa bb cc dd ee"
	for i := 0 ; i < b.N ; i++ {
		IsValidPassphrase(strings.Fields(passphrase))
	}
}

func TestPartTwoLineOne(t *testing.T) {
	passphrase := "abcde fghij"
	actual := IsOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestPartTwoLineTwo(t *testing.T) {
	passphrase := "abcde xyz ecdab"
	actual := IsOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != false {
		t.Error("Expected passphrase to be not valid:", passphrase)
	}
}

func TestPartTwoLineThree(t *testing.T) {
	passphrase := "a ab abc abd abf abj"
	actual := IsOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestPartTwoLineFour(t *testing.T) {
	passphrase := "iiii oiii ooii oooi oooo"
	actual := IsOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != true {
		t.Error("Expected passphrase to be valid:", passphrase)
	}
}

func TestPartTwoLineFive(t *testing.T) {
	passphrase := "oiii ioii iioi iiio"
	actual := IsOrderedValidPassphrase(strings.Fields(passphrase))
	if actual != false {
		t.Error("Expected passphrase to be not valid:", passphrase)
	}
}
