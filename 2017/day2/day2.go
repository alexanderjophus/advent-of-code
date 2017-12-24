package main

import (
	"math"
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func sumLine(input []int) int {
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

func dayTwo(fileName string) int {
	file, _ := os.Open(fileName)
	defer file.Close()
	checksum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list := []int{}
		for _, p := range strings.Fields(scanner.Text()) {
			i, _ := strconv.Atoi(p)
			list = append(list, i)
		}
		checksum += sumLine(list)
	}
	return checksum
}

func main()  {
	fmt.Println("Checksum =", dayTwo("input.txt"))
}