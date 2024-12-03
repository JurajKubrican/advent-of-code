package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func b() {
	scanner := getScanner("in.txt")
	reMuls := regexp.MustCompile(`(do|don't)\(\)|mul\((\d+),(\d+)\)`)
	reOperands := regexp.MustCompile(`\d+`)

	total := 0
	enable := true

	for scanner.Scan() {
		line := getLine(scanner)

		matches := findAllOccurrences(reMuls, line)

		fmt.Println(matches)

		for _, match := range matches {

			if match == "do()" {
				enable = true
			} else if match == `don't()` {
				enable = false
			} else if enable {
				operands := findAllOccurrences(reOperands, match)
				a, _ := strconv.Atoi(operands[0])
				b, _ := strconv.Atoi(operands[1])
				total += a * b
			}

		}
	}

	fmt.Println("DONE", total)

}
