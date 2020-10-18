package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day1Solve(inputFile string) {
	file, err := os.Open(inputFile)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "")

		p1 := day1(input)
		p2 := day1Bonus(input)
		fmt.Printf("d1p1 = %d\nd1p2 = %d\n", p1, p2)
	}
}

func day1(input []string) int {
	result := 0
	for index, value := range input {
		if input[(index+1)%len(input)] == value {
			intValue, _ := strconv.Atoi(value)
			result += intValue
		}
	}
	return result
}

func day1Bonus(input []string) int {
	result := 0
	half := len(input) / 2
	for i := 0; i < half; i++ {
		if input[i] == input[half+i] {
			intValue, _ := strconv.Atoi(input[i])
			result += intValue * 2
		}
	}
	return result
}
