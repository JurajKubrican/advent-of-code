package main

import (
	"fmt"
	"strconv"
	"strings"
)

func a() {
	scanner := getScanner("in.txt")

	total := 0

	rules := make([][]int, 0)
	// rules
	for scanner.Scan() {
		line := getLine(scanner)
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		rules = append(rules, []int{a, b})
		fmt.Println(rules)
	}

	for scanner.Scan() {
		line := getLine(scanner)
		nums := parseIntSlice(line, ",")
		fmt.Println(nums)

		pass := isValidLine(rules, nums)

		if pass {
			total += nums[len(nums)/2]
		}

		fmt.Println("PASS", pass)
	}

	fmt.Println("DONE", total)

}
func parseIntSlice(s string, separator string) []int {
	parts := strings.Split(s, separator)
	result := make([]int, len(parts))
	for i, p := range parts {
		res, err := strconv.Atoi(p)
		result[i] = res
		if err != nil {
			fmt.Println("Error converting to int", p)
		}
	}
	return result
}

func isValidLine(rules [][]int, line []int) bool {
	pass := true
	pastItems := make(map[int]struct{})
	for _, num := range line {
		// fmt.Println("Checking", num)
		pass = checkRules(rules, pastItems, num)
		if !pass {
			break
		}
		pastItems[num] = struct{}{}
	}
	return pass
}

func checkRules(rules [][]int, pastItems map[int]struct{}, num int) bool {
	for _, rule := range rules {
		if num == rule[0] { // current number is the second number in the rule
			// fmt.Println("FOUND", rule[0], rule[1], pastItems[rule[1]], pastItems)
			_, present := pastItems[rule[1]]
			if present {
				return false
			}
		}
	}
	return true
}
