package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func Test01(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}

	result := Solution(input)

	if result != 514579 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	// todo read a file and store into a slice
	f, err := os.Open("input.txt")

	if err != nil {
		t.Fail()
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		input = append(input, i)
	}

	result := Solution(input)

	if result != 0 {
		t.Fail()
	}
}
