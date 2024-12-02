package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func a() {
	//Read the file
	input, err := os.Open("test.txt")
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

		a, b, _ := parsePair(items[0], items[1])
		dir := spaceship(a, b)
		// fmt.Println(dir, a, b)
		lineRes := true
		// fmt.Println(items)

		for i := 0; i < len(items)-1; i++ {
			a, b, _ := parsePair(items[i], items[i+1])
			if dir == 0 {
				lineRes = false
				break
			}
			if dir == 1 {
				if !compare(a, b) {
					lineRes = false
					break
				}
			}
			if dir == -1 {
				if !compare(b, a) {
					lineRes = false
					break
				}
			}

		}
		fmt.Println(line, lineRes)
		result = append(result, lineRes)

	}

	trueCount := 0
	for _, res := range result {
		if res {
			trueCount++
		}
	}
	fmt.Println(trueCount)

}

func compare(a int, b int) bool {
	dif := b - a
	if dif <= 0 {
		return false
	}
	if dif > 3 {
		return false
	}
	return true
}

func parsePair(aStr string, bStr string) (int, int, error) {
	a, err := strconv.Atoi(aStr)
	if err != nil {
		fmt.Println("Error parsing integer:", err)
		return 0, 0, err
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		fmt.Println("Error parsing integer:", err)
		return 0, 0, err
	}

	return a, b, nil
}

// increasing = 1 decreasing = -1 equal = 0
func spaceship(a, b int) int {
	if a > b {
		return -1
	} else if a < b {
		return 1
	} else {
		return 0
	}
}
