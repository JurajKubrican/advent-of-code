package main

import (
	"fmt"
)

type Coords struct {
	X int
	Y int
}

func a() {
	scanner := getScanner("test.txt")

	total := 0

	antennas := make(map[rune][]Coords, 200)
	lastIndex := Coords{0, 0}
	j := -1
	for scanner.Scan() {
		j++
		line := getLine(scanner)
		lastIndex.Y = len(line) - 1
		for i, c := range line {

			if c == '.' {
				continue
			}
			if _, ok := antennas[c]; !ok {
				antennas[c] = make([]Coords, 0)
			}
			antennas[c] = append(antennas[c], Coords{j, i})
		}
	}
	lastIndex.X = j

	results := makeMatrix(lastIndex.X+1, lastIndex.Y+1, '.')
	fmt.Println(antennas)

	for sym, locations := range antennas {
		fmt.Println(string(sym), locations)

		for i := range locations {
			for j := i + 1; j < len(locations); j++ {
				a := locations[i]
				b := locations[j]
				antinodeA := getAntinode(a, b)
				if isIndexInBounds(antinodeA, lastIndex) {
					results[antinodeA.X][antinodeA.Y] = '#'
				}
				antinodeB := getAntinode(b, a)
				if isIndexInBounds(antinodeB, lastIndex) {
					results[antinodeB.X][antinodeB.Y] = '#'
				}

			}
		}
	}

	printMatrix(results)
	total = countInMatrix(results, '#')

	fmt.Println("DONE", total, lastIndex)

}

func getAntinode(a Coords, b Coords) Coords {
	dist := Coords{b.X - a.X, b.Y - a.Y}
	return Coords{b.X + dist.X, b.Y + dist.Y}
}

func makeMatrix(x, y int, seed rune) [][]byte {
	matrix := make([][]byte, x)
	for i := range matrix {
		matrix[i] = make([]byte, y)
		for j := range matrix[i] {
			matrix[i][j] = byte(seed)
		}
	}
	return matrix
}

func printMatrix(matrix [][]byte) {
	for i := range matrix {
		fmt.Println(string(matrix[i]))
	}
}
func isIndexInBounds(item Coords, lastIndex Coords) bool {
	return item.X >= 0 && item.X <= lastIndex.X && item.Y >= 0 && item.Y <= lastIndex.Y

}
func countInMatrix(matrix [][]byte, item rune) int {
	total := 0
	for i := range matrix {
		for j := range matrix[i] {
			if rune(matrix[i][j]) == item {
				total++
			}
		}
	}
	return total
}
