package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/JurajKubrican/advent-of-code/2023/files"
)

// dest source len
// 50 98 2

type Map struct {
	dest   int
	source int
	length int
}

// 50 98 2
func (rule Map) Process(num int) (bool, int) {
	if num < rule.source || num >= rule.source+rule.length {
		return false, 0
	}
	dist := num - rule.source

	return true, rule.dest + dist
}

func Day5A() {

	scanner := files.GetScanner("./day5/in.txt")

	sum := 0.0

	scanner.Scan()
	line := scanner.Text()
	seedArr := strings.Fields(strings.Split(line, ":")[1])

	fmt.Println(seedArr)

	hashmaps := make([][]Map, 0)

	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, "map") {
			hashmaps = append(hashmaps, make([]Map, 0))
			continue
		}

		row := strings.Fields(line)
		hashmap := hashmaps[len(hashmaps)-1]
		dest, _ := strconv.Atoi(row[0])
		source, _ := strconv.Atoi(row[1])
		length, _ := strconv.Atoi(row[2])
		hashmap = append(hashmap, Map{dest, source, length})
		hashmaps[len(hashmaps)-1] = hashmap
	}

	resultSeeds := make([]int, 0)
	for _, strSeed := range seedArr {
		seed, _ := strconv.Atoi(strSeed)
		fmt.Println("STARTING")
		for _, hashmap := range hashmaps {
			fmt.Println(seed)
			for _, hashmapRow := range hashmap {
				if ok, newSeed := hashmapRow.Process(seed); ok {
					fmt.Println("MATCH", seed, hashmapRow, newSeed)
					seed = newSeed
					break
				}
			}

		}
		resultSeeds = append(resultSeeds, seed)
	}

	fmt.Println(sum, resultSeeds)

	fmt.Println(slices.Min(resultSeeds))

}
