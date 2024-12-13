package main

import (
	"fmt"
	"strconv"
	"strings"
)

func a() {
	scanner := getScanner("in.txt")

	// a + 3
	// b + 1

	result := 0

	for scanner.Scan() {
		line := getLine(scanner)
		println(line)
		parts := strings.Split(line, "Button A: X+")
		parts = strings.Split(parts[1], ", Y+")
		ax, ay := parseIntArray(parts)

		scanner.Scan()
		line = getLine(scanner)
		parts = strings.Split(line, "Button B: X+")
		parts = strings.Split(parts[1], ", Y+")
		bx, by := parseIntArray(parts)

		scanner.Scan()
		line = getLine(scanner)
		parts = strings.Split(line, "Prize: X=")
		parts = strings.Split(parts[1], ", Y=")
		px, py := parseIntArray(parts)

		scanner.Scan()
		// COMPUTE

		res := compute(ax, bx, px, ay, by, py)
		if res > 0 {
			result += res
		}

		fmt.Println(res)

	}
	fmt.Println(result)
}

func parseIntArray(s []string) (int, int) {
	res := make([]int, len(s))
	for i, p := range s {
		r, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		res[i] = r
	}
	return res[0], res[1]
}

func compute(a, b, c, d, e, f int) int {
	y := (a*f - c*d) / (a*e - b*d)
	x := (c - b*y) / a

	println(x, y)
	if a*x+b*y != c || d*x+e*y != f {
		println("INVALID")
		return 0
	}

	return 3*x + y
}
