package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JurajKubrican/advent-of-code/2023/files"
)

func Day2A() {

	scanner := files.GetScanner("./day2/in.txt")

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		line, _ = strings.CutPrefix(line, "Game ")
		parts := strings.Split(line, ": ")

		id, _ := strconv.Atoi(parts[0])
		gameLines := strings.Split(parts[1], "; ")
		ok := true
		for _, gameLine := range gameLines {

			cubes := strings.Split(gameLine, ", ")
			for _, cubeLine := range cubes {
				cube := strings.Split(cubeLine, " ")
				numCubes, _ := strconv.Atoi(cube[0])
				fmt.Println(cubeLine)
				if cube[1] == "red" {
					if numCubes > 12 {
						ok = false
						break
					}
				} else if cube[1] == "green" {
					if numCubes > 13 {
						ok = false
						break
					}
				} else if cube[1] == "blue" {
					if numCubes > 14 {
						ok = false
						break
					}
				}
			}

		}
		fmt.Println(id, ok)
		if ok {
			sum += id
		}

	}

	fmt.Println(sum)

}
