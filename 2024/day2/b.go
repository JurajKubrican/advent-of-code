package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func b() {
	//Read the file
	input, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}

	scanner := bufio.NewScanner(input)

	// empty slice of strings that can be appended to
	var result []bool
	// line
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")

		lineRes, offender := evalLine(items)
		fmt.Println(1, items)
		if !lineRes {
			items = strings.Split(line, " ")
			shorterLine := append(items[:offender], items[offender+1:]...)

			fmt.Println(2, items)
			lineRes, _ = evalLine(shorterLine)
		}
		if !lineRes {
			items = strings.Split(line, " ")
			shorterLine := items[1:]
			lineRes, _ = evalLine(shorterLine)
		}
		if !lineRes {
			items = strings.Split(line, " ")
			shorterLine := append([]string{items[0]}, items[2:]...)
			lineRes, _ = evalLine(shorterLine)
		}

		result = append(result, lineRes)
		fmt.Println()
		fmt.Println()
	}

	fmt.Println(result)

	trueCount := 0
	for _, res := range result {
		if res {
			trueCount++
		}
	}
	fmt.Println(trueCount)

}

func evalLine(items []string) (bool, int) {
	firstA, firstB, _ := parsePair(items[0], items[1])
	dir := spaceship(firstA, firstB)
	lineRes := true
	i := 0

	fmt.Println(items)

	for i = 0; i < len(items)-1; i++ {
		a, b, _ := parsePair(items[i], items[i+1])
		if dir == 0 {
			lineRes = false
			fmt.Println(a, b, "dir 0")
			break
		}
		if dir == 1 && !compare(a, b) {
			lineRes = false
			fmt.Println(a, b, "dir 1")
			break
		}
		if dir == -1 && !compare(b, a) {
			lineRes = false
			fmt.Println(a, b, "dir -1")
			break
		}

	}

	fmt.Println()
	// i+1  is the offender
	return lineRes, i + 1
}
