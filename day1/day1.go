package day1

import (
	"strconv"
)

func DayOne(input []string) int {
	result := 0
	for index, value := range input {
		if input[(index+1)%len(input)] == value {
			intValue, _ := strconv.Atoi(value)
			result += intValue
		}
	}
	return result
}

func DayOneBonus(input []string) int {
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
