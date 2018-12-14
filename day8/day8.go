package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

}

func Day8(file *os.File) (int, error) {
	registers := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := strings.Split(scanner.Text(), " ")
		fmt.Printf("registration: %s\n", in[0])
		fmt.Printf(" instruction: %s\n", in[1])
		fmt.Printf("       value: %s\n", in[2])
		fmt.Printf(" conditional: %s %s %s\n", in[4], in[5], in[6])

		regValue, err := strconv.Atoi(in[0])
		if err != nil {
			return 0, err
		}
		value, err := strconv.Atoi(in[2])
		if err != nil {
			return 0, err
		}
		cond1, err := strconv.Atoi(in[4])
		if err != nil {
			return 0, err
		}
		cond2, err := strconv.Atoi(in[6])
		if err != nil {
			return 0, err
		}

		registers[in[0]] = calc(regValue, value, cond1, cond2, in[1], in[5])
	}

	max := 0
	largestRegister := ""
	for k, v := range registers {
		if max < v {
			max = v
			largestRegister = k
		}
	}

	return registers[largestRegister], nil
}

func calc(start, value, cond1, cond2 int, instruction, condition string) int {
	// var f = instructioner
	// f = add
	// return f(start, value)
	return 0
}

type instructioner interface {
}

// type instruction func(a int, b int) int

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}
