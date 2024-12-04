package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func b() {
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

		if isValidLine(rules, nums) {
			fmt.Println(nums, " - PASS")
			continue
		}

		sort.SliceStable(nums, func(a, b int) bool {
			for _, rule := range rules {
				if rule[0] == nums[a] && rule[1] == nums[b] { // current num should be on the left
					return true
				}
			}
			return false
		})

		if isValidLine(rules, nums) {
			fmt.Println(nums, " - SORTED")

		} else {
			fmt.Println(nums, " - UNSORTED!!!!!!")
		}

		fmt.Println("PASS", nums[len(nums)/2])
		total += nums[len(nums)/2]
	}

	fmt.Println("DONE", total)

}
