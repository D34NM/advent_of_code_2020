package main

import (
	"strings"
)

func getOuterBagColor(input string) string {
	index := strings.Index(input, "bag") - 1

	return input[:index]
}

func getInnerBagColors(input string) []string {
	parts := strings.Split(input, ",")

	result := make([]string, 0)
	for _, part := range parts {
		if strings.Index(part, "no other bag") > 0 {
			continue
		}

		index := strings.Index(part, "bag") - 1
		bag := part[3:index]
		result = append(result, bag)
	}

	return result
}

func traverse(key string, mapping map[string][]string) int {
	bags, ok := mapping[key]

	if ok {
		count := 0
		for _, bag := range bags {
			innerBags, _ := mapping[bag]
			if len(innerBags) > 0 {
				count = count + traverse(bag, mapping)
			} else {
				count = count + 1
			}
		}

		return count
	}

	return 1
}

func solutionOne(input []string) int {
	count := 0

	mapping := map[string][]string{}

	for _, line := range input {
		parts := strings.Split(line, "contain")

		outerBagColor := getOuterBagColor(parts[0])
		innerBagColors := getInnerBagColors(parts[1])

		for _, inner := range innerBagColors {
			val, ok := mapping[inner]
			if ok {
				val = append(val, outerBagColor)
				mapping[inner] = val
			} else {
				mapping[inner] = []string{outerBagColor}
			}
		}
	}

	bags, _ := mapping["shiny gold"]

	for _, bag := range bags {
		count = count + traverse(bag, mapping)
	}

	return count
}
