package main

import (
	"bufio"
	"os"
)

func main() {
	// a()
	b()

}

func getScanner(in string) *bufio.Scanner {

	input, _ := os.Open(in)

	return bufio.NewScanner(input)
}

func getLine(scanner *bufio.Scanner) string {
	line := scanner.Text()

	return line
}
