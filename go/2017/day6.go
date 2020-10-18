package main

import (
	"reflect"
	"strconv"
	"strings"
)

func Day6(input []int) int {
	iterations := 0
	seen := make(map[string]bool)
	for {
		if seen[String(input)] {
			return iterations
		}
		seen[String(input)] = true
		input = Day6Cycle(input)
		iterations++
	}
}

func Day6Bonus(input []int) int {
	hare := make([]int, len(input))
	tortoise := make([]int, len(input))
	copy(hare, input)
	copy(tortoise, input)
	for {
		hare = Day6Cycle(hare)
		tortoise = Day6Cycle(Day6Cycle(tortoise))
		if reflect.DeepEqual(hare, tortoise) {
			break
		}
	}
	iterations := 1
	for {
		hare = Day6Cycle(hare)
		if reflect.DeepEqual(hare, tortoise) {
			return iterations
		}
		iterations++
	}
}

func String(in []int) string {
	s := []string{}
	for k := range in {
		s = append(s, strconv.Itoa(in[k]))
	}
	return strings.Join(s, "-")
}

func Day6Cycle(input []int) []int {
	i, m := -1, -1
	for k, e := range input {
		if m < e {
			m = e
			i = k
		}
	}
	input[i] = 0
	for k := i; k < m+i; k++ {
		index := (k + 1) % len(input)
		input[index]++
	}
	return input
}
