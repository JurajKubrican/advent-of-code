package main

import (
	"fmt"
	"strconv"
	"strings"
)

func a() {
	scanner := getScanner("in.txt")

	scanner.Scan()
	line := scanner.Text()
	row := strings.Fields(line)

	// fmt.Println(row)
	for i := 0; i < 25; i++ {
		row = iterate(row)
		// fmt.Println(row)
	}

	fmt.Println("Total:", len(row))

}

func iterate(in []string) []string {
	out := make([]string, 0)
	for _, num := range in {
		if num == "0" {
			// fmt.Print("A", num)
			out = append(out, "1")

		} else if len(num)%2 == 0 {
			l := len(num) / 2
			a, _ := strconv.ParseInt(num[:l], 10, 64)
			out = append(out, strconv.FormatInt(a, 10))
			b, _ := strconv.ParseInt(num[l:], 10, 64)

			out = append(out, strconv.FormatInt(b, 10))
		} else {
			// fmt.Println("C", num)
			a, _ := strconv.ParseInt(num, 10, 64)
			a *= 2024
			out = append(out, strconv.FormatInt(a, 10))
		}
	}
	return out
}

func parseIntSlice(s string, separator string) []int {
	parts := strings.Split(s, separator)
	result := make([]int, len(parts))
	for i, p := range parts {
		res, err := strconv.Atoi(p)

		if err != nil {
			result[i] = -99
			fmt.Println("Error converting to int", p)
		} else {
			result[i] = res
		}
	}
	return result
}
