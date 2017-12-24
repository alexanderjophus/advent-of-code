package main

import (
	"math"
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func findRange(input []int) int {
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
	return high-low
}

func findAlt(input []int) int {
	for _, value := range input {
		for _, value2 := range input {
			if value > value2 && value%value2==0 {
				return value/value2
			}
		}
	}
	panic("Contract broken, not found one number divisible by another")
}

func dayTwo(fileName string) (int, int) {
	file, _ := os.Open(fileName)
	defer file.Close()
	rangeAddition := 0
	divisorAddition := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list := []int{}
		for _, p := range strings.Fields(scanner.Text()) {
			i, _ := strconv.Atoi(p)
			list = append(list, i)
		}
		rangeAddition += findRange(list)
		divisorAddition += findAlt(list)
	}
	return rangeAddition, divisorAddition
}

func main()  {
	rangeAddition, divisorAddition := dayTwo("input.txt")
	fmt.Println("Checksum =", rangeAddition)
	fmt.Println("Checksum2 =", divisorAddition)
}