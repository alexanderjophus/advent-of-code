package main

import (
	"fmt"
	"./solution"
	"os"
)


func main()  {
	file, err := os.Open("2017/day2/go/input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	rangeAddition, divisorAddition := solution.DayTwo(file)
	fmt.Println("Checksum =", rangeAddition)
	fmt.Println("Checksum2 =", divisorAddition)
}