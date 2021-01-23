package main

import (
	"bufio"
	"os"
	"testing"
)

func Test01(t *testing.T) {
	input := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}

	result := solution(input)

	if result != 2 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	f, err := os.Open("input.txt")

	if err != nil {
		t.Fail()
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	input = append(input, "")
	result := solution(input)

	if result != 230 {
		t.Fail()
	}
}
