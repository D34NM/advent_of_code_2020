package main

import (
	"strings"
)

func getOuterBagColor(input string) string {
	index := strings.Index(input, "bag") - 1

	return input[:index]
}

func getInnerBagColors(input string) *[]string {
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

	return &result
}

var seen = make(map[string]int, 0)

func traverse(key string, mapping *map[string]*[]string) int {
	if _, exists := seen[key]; exists {
		return 0
	}

	if bags, exists := (*mapping)[key]; exists {
		count := 0
		for _, bag := range *bags {
			if _, exists := (*mapping)[bag]; exists {
				count = count + traverse(bag, mapping)
			} else {
				count++
			}
		}

		seen[key] = count
		return count
	}

	seen[key] = 1
	return 1
}

func solutionOne(input *[]string) int {
	count := 0
	mapping := map[string]*[]string{}

	for _, line := range *input {
		parts := strings.Split(line, "contain")

		outerBagColor := getOuterBagColor(parts[0])
		innerBagColors := getInnerBagColors(parts[1])

		for _, innerBag := range *innerBagColors {
			if val, exists := mapping[innerBag]; exists {
				*val = append(*val, outerBagColor)
				mapping[innerBag] = val
			} else {
				mapping[innerBag] = &[]string{outerBagColor}
			}
		}
	}

	bags, _ := mapping["shiny gold"]

	for _, bag := range *bags {
		count = count + traverse(bag, &mapping)
	}

	return count
}
