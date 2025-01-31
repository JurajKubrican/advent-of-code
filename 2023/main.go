package main

import (
	"bufio"
	"os"

	"github.com/JurajKubrican/advent-of-code/2023/day2"
)

func main() {

	// day2.Day2A()
	day2.Day2B()
}

func getScanner(in string) *bufio.Scanner {

	input, _ := os.Open(in)

	return bufio.NewScanner(input)
}
