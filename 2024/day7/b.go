package main

import (
	"fmt"
	"strconv"
	"strings"
)

func b() {
	scanner := getScanner("in.txt")

	total := 0

	for scanner.Scan() {
		line := getLine(scanner)
		parts := strings.Split(line, ":")
		res, _ := strconv.Atoi(parts[0])
		operands, _ := convertSliceToInt(strings.Fields(parts[1]))

		fmt.Println(res, operands)

		success := computeB(0, operands, res)
		if success {
			total += res
		}

		fmt.Println(res, operands, success)
	}

	fmt.Println("DONE", total)

}

func computeB(a int, operands []int, res int) bool {
	if a > res {
		return false
	}
	if len(operands) == 0 {
		if a == res {
			return true
		}
		return false
	}

	b := operands[0]
	if len(operands) > 0 {
		if computeB(a+b, operands[1:], res) {
			return true
		}
		if computeB(a*b, operands[1:], res) {
			return true
		}
		if computeB(concat(a, b), operands[1:], res) {
			return true
		}
	}

	return false
}

func concat(a int, b int) int {

	res, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		panic(err)
	}
	return res
}
