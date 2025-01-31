package day2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/JurajKubrican/advent-of-code/2023/files"
)

func Day2B() {

	scanner := files.GetScanner("./day2/in.txt")

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		line, _ = strings.CutPrefix(line, "Game ")
		parts := strings.Split(line, ": ")

		id, _ := strconv.Atoi(parts[0])
		gameLines := strings.Split(parts[1], "; ")

		cubeNumbers := map[string]int{}
		for _, gameLine := range gameLines {

			cubes := strings.Split(gameLine, ", ")
			for _, cubeLine := range cubes {
				cube := strings.Split(cubeLine, " ")
				numCubes, _ := strconv.Atoi(cube[0])

				cubeColor := cube[1]
				num := cubeNumbers[cubeColor]
				cubeNumbers[cubeColor] = slices.Max([]int{num, numCubes})
			}

		}
		fmt.Println(id, cubeNumbers)
		power := 1
		for _, v := range cubeNumbers {
			power *= v
		}
		sum += power

	}

	fmt.Println(sum)

}
