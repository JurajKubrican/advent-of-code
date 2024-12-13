package main

import "fmt"

type G struct {
	id      int
	corners int
	area    int
}

func b() {
	scanner := getScanner("in.txt")

	matrix := make([][]Point, 0)
	i := 0
	for scanner.Scan() {
		line := getLine(scanner)
		matrix = append(matrix, make([]Point, len(line)))
		for j, c := range line {
			matrix[i][j] = Point{c: c, perimeter: 0}
		}
		i++
	}

	id = 0
	groups = map[int]G{}
	for i, row := range matrix {
		for j, c := range row {
			id++
			area := compute2(matrix, i, j, c, id)
			groups[id] = G{id: id, area: area}

		}
	}

	for i, row := range matrix {
		for j := range row {
			id++
			patch := getPatchCopy(matrix, i, j, Point{})
			for k := 0; k < 4; k++ {
				isCorner := isCorner(patch)
				isSuperCorner := isSuperCorner(patch)

				if isCorner {
					fmt.Println("CORNER", i, j)
					id := patch[1][1].id
					g := groups[id]
					g.corners++
					groups[id] = g
				}
				if isSuperCorner {
					fmt.Println("SUPER CORNER", i, j)
					id := patch[0][0].id
					g := groups[id]
					g.corners++
					groups[id] = g
				}

				patch = rotatePatch(patch)
			}

		}
	}

	printM(matrix)
	println()

	printI(matrix)

	res := 0
	for _, g := range groups {
		fmt.Println(g)
		res += g.area * g.corners
	}

	fmt.Println(res)

}

var id int

var groups map[int]G

// var re

func compute2(matrix [][]Point, i, j int, g Point, id int) int {
	if matrix[i][j].c != g.c {
		return 0
	}
	if matrix[i][j].visited == true {
		return 0
	}
	// println(i, j, string(matrix[i][j].c))

	matrix[i][j].visited = true
	matrix[i][j].id = id

	res := 1

	if isIndexInBounds(matrix, i-1, j) {
		res += compute2(matrix, i-1, j, g, id)
	}
	if isIndexInBounds(matrix, i+1, j) {
		res += compute2(matrix, i+1, j, g, id)
	}
	if isIndexInBounds(matrix, i, j-1) {
		res += compute2(matrix, i, j-1, g, id)
	}
	if isIndexInBounds(matrix, i, j+1) {
		res += compute2(matrix, i, j+1, g, id)
	}
	return res

}

func rotatePatch(patch [][]Point) [][]Point {
	n := len(patch)
	res := make([][]Point, n)
	for i := 0; i < n; i++ {
		res[i] = make([]Point, n)
		for j := 0; j < n; j++ {
			res[i][j] = patch[j][n-1-i]
		}
	}
	return res
}

func printI(matrix [][]Point) {
	for _, row := range matrix {
		for _, c := range row {
			print(" ", c.id)
		}
		println()
	}
}

func isCorner(patch [][]Point) bool {
	p := patch[1][1]

	b := patch[0][1]
	c := patch[1][0]

	if p.c != b.c && p.c != c.c {
		return true
	}
	return false
}

func isSuperCorner(patch [][]Point) bool {

	a := patch[0][0]
	b := patch[0][1]
	c := patch[1][0]

	if isCorner(patch) && a.c == b.c && a.c == b.c && a.c == c.c {
		return true
	}

	return false
}
