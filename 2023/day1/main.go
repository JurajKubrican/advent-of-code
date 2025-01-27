package main

import (
	"bufio"
	"os"
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
