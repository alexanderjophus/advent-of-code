package io

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFileIntSlice(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ret []int
	for scanner.Scan() {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ret = append(ret, a)
	}
	return ret, nil
}
