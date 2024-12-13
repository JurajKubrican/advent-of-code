package main

import (
	"fmt"
	"strings"
)

func b() {
	scanner := getScanner("in.txt")

	// a + 3
	// b + 1

	result := int64(0)

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

		res := compute2(int64(ax), int64(bx), 10000000000000+int64(px), int64(ay), int64(by), 10000000000000+int64(py))
		if res > 0 {
			result += res
		}

		fmt.Println(res)

	}
	fmt.Println(result)
}

func compute2(a, b, c, d, e, f int64) int64 {
	y := (a*f - c*d) / (a*e - b*d)
	x := (c - b*y) / a

	println(x, y)
	if a*x+b*y != c || d*x+e*y != f {
		println("INVALID")
		return 0
	}

	return 3*x + y
}
