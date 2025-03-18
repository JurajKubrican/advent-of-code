package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/JurajKubrican/advent-of-code/2023/files"
)

func Day5B() {

	scanner := files.GetScanner("./day5/in.txt")

	sum := 0.0

	scanner.Scan()
	line := scanner.Text()
	sourceSeedArr := strings.Fields(strings.Split(line, ":")[1])

	seedArr := make([]int, 0)
	for i := 0; i < len(sourceSeedArr); i += 2 {
		from, _ := strconv.Atoi(sourceSeedArr[0])
		length, _ := strconv.Atoi(sourceSeedArr[1])
		for j := from; j < from+length; j++ {
			seedArr = append(seedArr, j)
		}
	}

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
	for _, seed := range seedArr {
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
