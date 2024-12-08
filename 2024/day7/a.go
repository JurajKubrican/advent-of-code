package main

import (
	"fmt"
	"strconv"
	"strings"
)

func a() {
	scanner := getScanner("in.txt")

	total := 0

	for scanner.Scan() {
		line := getLine(scanner)
		parts := strings.Split(line, ":")
		res, _ := strconv.Atoi(parts[0])
		operands, _ := convertSliceToInt(strings.Fields(parts[1]))

		fmt.Println(res, operands)

		success := compute(res, operands)
		if success {
			total += res
		}

		fmt.Println(res, operands, success)
	}

	fmt.Println("DONE", total)

}

func compute(res int, operands []int) bool {
	if len(operands) == 1 {
		return false
	}
	cur := operands[len(operands)-1]
	if res%cur == 0 {
		if res/cur == operands[0] {
			return true
		}
		// fmt.Println("delving into *", res/cur, operands[:len(operands)-1])
		if compute(res/cur, operands[:len(operands)-1]) {
			return true
		}
	}

	if res-cur > 0 {
		if res-cur == operands[0] {
			return true
		}
		// fmt.Println("delving into +", res-cur, operands[:len(operands)-1])
		if compute(res-cur, operands[:len(operands)-1]) {
			return true
		}
	}

	return false
}

func convertSliceToInt(strSlice []string) ([]int, error) {
	intSlice := make([]int, len(strSlice))
	for i, str := range strSlice {
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting to int", str)
			return nil, err
		}
		intSlice[i] = val
	}
	return intSlice, nil
}
