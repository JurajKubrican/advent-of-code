package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Read the file
	input, _ := os.Open("in.txt")

	scanner := bufio.NewScanner(input)

	globalMax := 0
	localMax := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			globalMax = max(globalMax, localMax)

			localMax = 0
		} else {
			item, _ := strconv.Atoi(line)
			localMax += item
		}

	}

	fmt.Println(globalMax)

}
