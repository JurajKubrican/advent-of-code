package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	//Read the file
	input, _ := os.Open("in.txt")

	scanner := bufio.NewScanner(input)

	globalMax := []int{}
	localMax := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			globalMax = append(globalMax, localMax)

			localMax = 0
		} else {
			item, _ := strconv.Atoi(line)
			localMax += item
		}

	}

	sort.Sort(sort.Reverse(sort.IntSlice(globalMax)))

	fmt.Println(globalMax[0] + globalMax[1] + globalMax[2])

}
