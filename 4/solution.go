package main

import (
	"regexp"
	"strconv"
	"strings"
)

func main() {}

func validBirthYear(byr string) bool {
	year, err := strconv.Atoi(byr)

	if err != nil {
		return false
	}

	if year >= 1920 && year <= 2002 {
		return true
	}

	return false
}

func validIssueYear(iyr string) bool {
	year, err := strconv.Atoi(iyr)

	if err != nil {
		return false
	}

	if year >= 2010 && year <= 2020 {
		return true
	}

	return false
}

func validExpirationYear(eyr string) bool {
	year, err := strconv.Atoi(eyr)

	if err != nil {
		return false
	}

	if year >= 2020 && year <= 2030 {
		return true
	}

	return false
}

func validHeight(hgt string) bool {
	cmIndex := strings.Index(hgt, "cm")
	inIndex := strings.Index(hgt, "in")

	if cmIndex > 0 {
		height, err := strconv.Atoi(hgt[:cmIndex])
		if err != nil {
			return false
		}
		if height >= 150 && height <= 193 {
			return true
		}
	}

	if inIndex > 0 {
		height, err := strconv.Atoi(hgt[:inIndex])
		if err != nil {
			return false
		}

		if height >= 59 && height <= 76 {
			return true
		}
	}

	return false
}

func validHairColour(ecl string) bool {
	r, _ := regexp.Compile("^#[a-f0-9]{6}$")

	if r.MatchString(ecl) {
		return true
	}

	return false
}

func validEyeColour(ecl string) bool {
	m := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

	if m[ecl] {
		return true
	}

	return false
}

func validPassportID(pid string) bool {
	r, _ := regexp.Compile("[0-9]{9}")

	if r.MatchString(pid) {
		return true
	}

	return false
}

func validCountryID(_ string) bool {
	return false
}

func solution(input []string) int {
	valid := 0
	empty := ""
	required := map[string]func(string) bool{
		"byr": validBirthYear,
		"iyr": validIssueYear,
		"eyr": validExpirationYear,
		"hgt": validHeight,
		"hcl": validHairColour,
		"ecl": validEyeColour,
		"pid": validPassportID,
		"cid": validCountryID,
	}

	current := 0
	for _, row := range input {
		if row == empty {
			if current >= len(required)-1 {
				valid++
			}
			current = 0
			continue
		}

		fields := strings.Fields(row)
		for _, field := range fields {
			quoteIndex := strings.Index(field, ":")

			f, _ := required[field[:quoteIndex]]
			valueIndex := quoteIndex + 1
			if f(field[valueIndex:]) {
				current++
			}
		}
	}

	return valid
}
