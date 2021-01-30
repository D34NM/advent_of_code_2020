package main

import "testing"

func Test01(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	result := solutionOne(input)

	if result != 5 {
		t.Fail()
	}
}

func Test02(t *testing.T) {

}
