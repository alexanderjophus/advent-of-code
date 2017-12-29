package main

import (
	"fmt"
	"./solution"
	"os"
)

func main()  {
	file, err := os.Open("2017/day5/go/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println("Steps part 1: ", solution.Day5(solution.ReadFile("2017/day5/go/input.txt")))
	fmt.Println("Steps part 2: ", solution.Day5Bonus(solution.ReadFile("2017/day5/go/input.txt")))
}

