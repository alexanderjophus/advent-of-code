package main

import (
	twentyseventeen "github.com/trelore/advent-of-code/go/2017"
)

func main() {
	twentyseventeen.Day1Solve("2017/day1.txt")
	twentyseventeen.Day2Solve("2017/day2.txt")
	twentyseventeen.Day3Solve("2017/day3.txt")
	twentyseventeen.Day4Solve("2017/day4.txt")
	twentyseventeen.Day5Solve("2017/day5.txt")
	twentyseventeen.Day6Solve("2017/day6.txt")
	// get original input -> put into file and conform to new standard
	// day6.Day6([]int{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6})
	// day6.Day6Bonus([]int{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6})
	twentyseventeen.Day7Solve("2017/day7.txt")
	twentyseventeen.Day8Solve("2017/day8.txt")
}
