package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func day5(input []int) int {
	counter := 0
	for i := 0 ; 0 <= i && i < len(input) ; counter++ {
		input[i] ++
		i = i + input[i] - 1
	}
	return counter
}

func day5Bonus(input []int) int {
	counter := 0
	for i := 0 ; 0 <= i && i < len(input) ; counter++ {
		if input[i] >= 3 {
			input[i] --
			i = i + input[i] + 1
		} else {
			input[i] ++
			i = i + input[i] - 1
		}
	}
	return counter
}

func readFile(filename string) []int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ret []int
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		ret = append(ret, a)
	}
	return ret
}

func main()  {
	fmt.Println("Steps part 1: ", day5(readFile("input.txt")))
	fmt.Println("Steps part 2: ", day5Bonus(readFile("input.txt")))
}
