package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read the file
	// a()
	b()

}

func getScanner(in string) *bufio.Scanner {

	input, _ := os.Open(in)

	return bufio.NewScanner(input)
}

func getLineOfInts(scanner *bufio.Scanner) []int {
	line := scanner.Text()
	parts := strings.Fields(line)
	ints := make([]int, len(parts))

	// Convert each part to an integer
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting to int", part)
			return nil
		}
		ints[i] = num
	}

	return ints
}
