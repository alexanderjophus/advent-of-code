package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func day4Solve(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//p1 consuming file?
	p1 := CountValidPassphrases(file, IsValidPassphrase)
	file.Seek(0, 0)
	p2 := CountValidPassphrases(file, IsOrderedValidPassphrase)
	fmt.Printf("d4p1 = %d\nd4p2 = %d\n", p1, p2)
}

func CountValidPassphrases(file *os.File, f func([]string) bool) int {
	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var list []string
		for _, p := range strings.Fields(scanner.Text()) {
			list = append(list, p)
		}
		if f(list) {
			counter++
		}
	}
	return counter
}

func IsValidPassphrase(input []string) bool {
	set := StringSet{}
	for _, value := range input {
		if set.add(value) == false {
			return false
		}
	}
	return true
}

func IsOrderedValidPassphrase(input []string) bool {
	set := StringSet{}
	for _, value := range input {
		if set.addOrdered(value) == false {
			return false
		}
	}
	return true
}

type StringSet struct {
	set map[string]bool
}

func (set *StringSet) add(i string) bool {
	if set.set == nil {
		set.set = make(map[string]bool)
	}
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *StringSet) addOrdered(i string) bool {
	tmp := strings.Split(i, "")
	sort.Strings(tmp)
	i = strings.Join(tmp, "")
	if set.set == nil {
		set.set = make(map[string]bool)
	}
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}
