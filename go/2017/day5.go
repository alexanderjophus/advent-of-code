package twentyseventeen

// import (
// 	"fmt"

// 	"github.com/trelore/advent-of-code/go/io"
// )

// func Day5Solve(inputFile string) {
// 	input, err := io.ReadFileIntSlice(inputFile)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	p1 := day5(input)
// 	p2 := day5Bonus(input)
// 	fmt.Printf("d5p1 = %d\nd5p2 = %d\n", p1, p2)
// }

// func day5(input []int) int {
// 	counter := 0
// 	for i := 0; 0 <= i && i < len(input); counter++ {
// 		input[i]++
// 		i = i + input[i] - 1
// 	}
// 	return counter
// }

// func day5Bonus(input []int) int {
// 	counter := 0
// 	for i := 0; 0 <= i && i < len(input); counter++ {
// 		if input[i] >= 3 {
// 			input[i]--
// 			i = i + input[i] + 1
// 		} else {
// 			input[i]++
// 			i = i + input[i] - 1
// 		}
// 	}
// 	return counter
// }
