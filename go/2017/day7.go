package twentyseventeen

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Day7Solve(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	p1 := day7(file)
	// p2 := Day3Bonus(input)
	fmt.Printf("d7p1 = %s\n", p1)
}

var r = regexp.MustCompile(`[a-zA-Z]+`)
var groupNames = r.SubexpNames()

func day7(file *os.File) string {
	stacks := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		items := r.FindAllStringSubmatch(text, -1)
		for _, v := range items {
			stacks[v[0]]++
		}
	}

	for k, v := range stacks {
		if v == 1 {
			return k
		}
	}

	return ""
}
