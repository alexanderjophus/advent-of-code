package twentyseventeen

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day2Solve(inputFile string) {
	file, err := os.Open(inputFile)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	p1, p2 := day2(file)
	fmt.Printf("d2p1 = %d\nd2p2 = %d\n", p1, p2)
}

func day2(file *os.File) (int, int) {
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

// FindRange of a slice of numbers
//
// e.g. [1,3,2,7] returns 6
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

// FindAlt returns the first value where one number is divisible by another with no remainder
//
// e.g. [3, 8, 6, 5] returns 2, as the 6 is divisible by 3
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
