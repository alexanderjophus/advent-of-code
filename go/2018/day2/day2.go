package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve(inputFile string) {
	file, err := os.Open("day2/input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	p1, p2 := DayTwo(file)
	fmt.Printf("d2p1 = %d\nd2p2 = %d\n", p1, p2)
}

func FindRange(input []int) int {
	low := math.MaxInt64
	high := 0
	for _, value := range input {
		if low > value {
			low = value
		}
		if high < value {
			high = value
		}
	}
	return high - low
}

func FindAlt(input []int) int {
	for _, value := range input {
		for _, value2 := range input {
			if value > value2 && value%value2 == 0 {
				return value / value2
			}
		}
	}
	panic("Contract broken, not found one number divisible by another")
}

func DayTwo(file *os.File) (int, int) {
	rangeAddition := 0
	divisorAddition := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list := []int{}
		for _, p := range strings.Fields(scanner.Text()) {
			i, _ := strconv.Atoi(p)
			list = append(list, i)
		}
		rangeAddition += FindRange(list)
		divisorAddition += FindAlt(list)
	}
	return rangeAddition, divisorAddition
}
