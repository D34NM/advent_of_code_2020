package main

import "strings"

func solutionOne(input []string) int {
	valid := 0
	empty := ""
	reserved := map[string]bool{"byr": true, "iyr": true, "eyr": true, "hgt": true, "hcl": true, "ecl": true, "pid": true}

	current := 0
	for _, row := range input {
		if row == empty {
			if current >= len(reserved) {
				valid++
			}
			current = 0
			continue
		}

		fields := strings.Fields(row)
		for _, field := range fields {
			quoteIndex := strings.Index(field, ":")

			if reserved[field[:quoteIndex]] {
				current++
			}
		}
	}

	return valid
}
