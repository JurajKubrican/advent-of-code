package main

import (
	"fmt"
	"strconv"
	"strings"
)

var cache = map[string]int{}

func b() {
	scanner := getScanner("in.txt")

	scanner.Scan()
	line := scanner.Text()
	rowStr := strings.Fields(line)

	res := iterate2(rowStr, 75)

	fmt.Println("Total:", res)

}

func iterate2(in []string, i int) int {
	key := getCacheKey(in, i)
	if res, found := cache[key]; found {
		return res
	}
	if i == 0 {
		return len(in)
	}

	res := 0
	for _, num := range in {
		nextStep := iterate([]string{num})
		res += iterate2(nextStep, i-1)
	}

	cache[key] = res
	return res
}

func getCacheKey(in []string, i int) string {
	res := strconv.Itoa(i)
	res += "-"
	res += strings.Join(in, " ")
	return res
}
