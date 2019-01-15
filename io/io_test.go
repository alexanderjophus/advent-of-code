package io_test

import (
	"reflect"
	"testing"

	"github.com/trelore/advent-of-code/io"
)

func TestFromFile(t *testing.T) {
	nums, err := io.ReadFileIntSlice("test_num.txt")
	if err != nil {
		t.Error(err)
	}
	expected := []int{0, 3, 0, 1, -3}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("got %v, expected %v", nums, expected)
	}
}
