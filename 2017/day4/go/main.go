package main

import (
	"fmt"
	"./solution"
	"os"
)

func main()  {
	file, err := os.Open("2017/day4/go/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println("Part 1:", solution.CountValidPassphrases(file, solution.IsValidPassphrase))
	fmt.Println("Part 2:", solution.CountValidPassphrases(file, solution.IsOrderedValidPassphrase))
}

