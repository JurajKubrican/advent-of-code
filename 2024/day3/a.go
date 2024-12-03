package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func a() {
	scanner := getScanner("in.txt")
	reMuls := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	reOperands := regexp.MustCompile(`\d+`)

	total := 0

	for scanner.Scan() {
		line := getLine(scanner)

		matches := findAllOccurrences(reMuls, line)

		for _, match := range matches {

			operands := findAllOccurrences(reOperands, match)
			fmt.Println(operands)
			a, _ := strconv.Atoi(operands[0])
			b, _ := strconv.Atoi(operands[1])
			total += a * b
		}
	}

	fmt.Println("DONE", total)

}

func findAllOccurrences(re *regexp.Regexp, text string) []string {
	matches := re.FindAllString(text, -1)
	return matches
}
