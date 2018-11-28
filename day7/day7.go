package day7

import (
	"bufio"
	"os"
	"regexp"
)

var r = regexp.MustCompile(`[a-zA-Z]+`)
var groupNames = r.SubexpNames()

func Day7(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

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
			return k, nil
		}
	}

	return "", nil
}
