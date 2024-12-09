package main

import (
	"fmt"
)

func b() {
	scanner := getScanner("in.txt")

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
				results[a.X][a.Y] = '#'
				antinodesA := getAntinodesInBounds(a, b, lastIndex)
				fmt.Println("A", antinodesA)
				for _, antinodeA := range antinodesA {
					if antinodeA.X == 10 && antinodeA.Y == 0 {
						fmt.Println("THIS FUCKER a,b", a, b, "NODES: ", antinodesA)
					}
					results[antinodeA.X][antinodeA.Y] = '#'
				}

				results[b.X][b.Y] = '#'
				antinodesB := getAntinodesInBounds(b, a, lastIndex)
				fmt.Println("B", antinodesB)
				for _, antinodeB := range antinodesB {
					results[antinodeB.X][antinodeB.Y] = '#'
				}

			}
		}
	}

	printMatrix(results)
	total = countInMatrix(results, '#')

	fmt.Println("DONE", total, lastIndex)

}

func getAntinodesInBounds(a Coords, b Coords, lastIndex Coords) []Coords {
	res := make([]Coords, 0)
	dist := Coords{intAbs(b.X - a.X), intAbs(b.Y - a.Y)}
	dir := Coords{sign(b.X - a.X), sign(b.Y - a.Y)}
	for i := 1; ; i++ {
		antinode := Coords{b.X + i*dist.X*dir.X, b.Y + i*dist.Y*dir.Y}
		if !isIndexInBounds(antinode, lastIndex) {
			break
		}
		res = append(res, antinode)
	}
	return res
}

func sign(x int) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
