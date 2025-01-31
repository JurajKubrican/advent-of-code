package files

import (
	"bufio"
	"os"
)

func GetScanner(in string) *bufio.Scanner {

	input, err := os.Open(in)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(input)
}
