package day3

import "math"

// Day3 Algorithm relies on square numbers. Designed for efficiency this algorithm avoids costly operations such as sqrt
// First, determine the floor of the root of the number-1. i.e. 12 -> 3, 21 -> 4, 1 -> 0
// Second, determine the numbers between i**2 and (i+1)**2 with the smallest manhattan distance,
// i.e. 1, 2, 11, 28 horizontally and 1, 4, 15, 34 vertically
// Third, find the minimum between the number and the v/h numbers and add that to the distance to the h/v number
func Day3(goal float64) int {
	i := 0.0
	for ; math.Pow(i+1, 2) < goal; i++ {
	}

	v := math.Pow(i, 2) + math.Ceil(i/2)
	h := math.Pow(i+1, 2) - math.Floor(i/2)

	return int(math.Min(math.Abs(goal-v), math.Abs(goal-h)) + math.Ceil(i/2))
}

func Day3Bonus(goal int) int {
	panic("not implemented yet")
}
