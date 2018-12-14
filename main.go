package main

import (
	"github.com/alexanderjosephtrelore/advent-of-code/day1"
	"github.com/alexanderjosephtrelore/advent-of-code/day2"
	"github.com/alexanderjosephtrelore/advent-of-code/day3"
	"github.com/alexanderjosephtrelore/advent-of-code/day4"
	"github.com/alexanderjosephtrelore/advent-of-code/day5"
	"github.com/alexanderjosephtrelore/advent-of-code/day7"
	"github.com/alexanderjosephtrelore/advent-of-code/day8"
)

var (
	day  string
	part string
)

func main() {
	day1.Solve("day1/input.txt")
	day2.Solve("day2/input.txt")
	day3.Solve("day3/input.txt")
	day4.Solve("day4/input.txt")
	day5.Solve("day5/input.txt")
	// get original input -> put into file and conform to new standard
	// day6.Day6([]int{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6})
	// day6.Day6Bonus([]int{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6})
	day7.Solve("day7/input.txt")
	day8.Solve("day8/input.txt")
}
