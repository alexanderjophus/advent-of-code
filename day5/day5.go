package day5

func Day5(input []int) int {
	counter := 0
	for i := 0; 0 <= i && i < len(input); counter++ {
		input[i]++
		i = i + input[i] - 1
	}
	return counter
}

func Day5Bonus(input []int) int {
	counter := 0
	for i := 0; 0 <= i && i < len(input); counter++ {
		if input[i] >= 3 {
			input[i]--
			i = i + input[i] + 1
		} else {
			input[i]++
			i = i + input[i] - 1
		}
	}
	return counter
}
