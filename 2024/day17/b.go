package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fnb() {
	scanner := getScanner("in.txt")

	scanner.Scan()
	line := getLine(scanner)
	// line = strings.Split(line, "Register A: ")[1]
	// __, _ := strconv.Atoi(line)

	scanner.Scan()
	line = getLine(scanner)
	line = strings.Split(line, "Register B: ")[1]
	initialB, _ := strconv.Atoi(line)

	scanner.Scan()
	line = getLine(scanner)
	line = strings.Split(line, "Register C: ")[1]
	initialC, _ := strconv.Atoi(line)

	scanner.Scan()
	scanner.Scan()
	line = getLine(scanner)
	sprogram := strings.Split(strings.Split(line, "Program: ")[1], ",")

	for _, v := range sprogram {
		i, _ := strconv.Atoi(v)
		program = append(program, i)
	}

	for i := 9999999999999999; ; i-- {
		A = 0
		B = initialB
		C = initialC
		output = []int{}

		if i%100000000 == 0 {
			fmt.Println(i)
		}

		if runProgram2() {
			return
		}

	}

}

func runProgram2() bool {
	lastLen := 0
	for index < len(program) {
		instructionMap[program[index]]()
		outLen := len(output)
		if outLen > lastLen {
			lastLen = outLen

			if output[outLen-1] != program[outLen-1] {
				return false
			}
			if outLen > len(program) {
				return false
			}
			if len(output) == len(program) {
				fmt.Println(A)
				return true
			}
		}

	}
	return false
}
