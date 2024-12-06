package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func b() {
	inputFile := "in.txt"
	if len(os.Args) == 2 {
		inputFile = os.Args[1]
	}

	scanner := getScanner(inputFile)
	total := 0

	// LOAD MATRIX
	matrix := make([][]rune, 0)
	guard := make([]int, 2)

	for scanner.Scan() {
		line := getLine(scanner)
		parts := []rune(line)

		index := slices.Index(parts, '^')
		if index != -1 {
			guard[0] = len(matrix)
			guard[1] = index
		}
		matrix = append(matrix, parts)
	}

	prevChar := ' '
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '#' {
				continue
			}

			// tempMatrix := deepCopyMatrix(matrix)
			tempGuard := make([]int, 2)
			copy(tempGuard, guard)
			tempGuardDir := '^'
			prevChar = matrix[i][j]
			matrix[i][j] = '#'

			result := traverseIsLoop(matrix, tempGuard, tempGuardDir)
			if result {
				total++
			}
			matrix[i][j] = prevChar
		}
	}

	fmt.Println(total)
	// printMatrix(matrix)

}

func traverseIsLoop(matrix [][]rune, guard []int, guardDir rune) bool {
	hitMap := make([][][]int, len(matrix))

	for true {
		next := getNeighbor(matrix, guard[0], guard[1], guardDir)
		if next == '~' {
			return false
		}
		if next == '#' {

			guardDir = nextDirection(guardDir)
			res := countHitFromDirection(hitMap, guard[0], guard[1], guardDir)
			if res == true {
				return true
			}
			// fmt.Println("TURN", guardDir)
			continue
		}
		if next == '.' || next == 'X' || next == '^' {
			guard[0], guard[1] = proceed(guard[0], guard[1], guardDir)
			// fmt.Println("PROCEED", guardDir)
			matrix[guard[0]][guard[1]] = 'X'
			continue
		}

	}
	return false
}

func deepCopyMatrix(matrix [][]rune) [][]rune {
	res := make([][]rune, len(matrix))
	for i := range matrix {
		res[i] = make([]rune, len(matrix[i]))
		copy(res[i], matrix[i])
	}
	return res
}

func countHitFromDirection(hitMap [][][]int, i, j int, dir rune) bool {
	if hitMap[i] == nil {
		hitMap[i] = make([][]int, len(hitMap))
	}
	if hitMap[i][j] == nil {
		hitMap[i][j] = make([]int, 4)
	}
	res := 0
	switch dir {
	case '^':
		hitMap[i][j][0]++
		res = hitMap[i][j][0]
	case '>':
		hitMap[i][j][1]++
		res = hitMap[i][j][1]
	case 'v':
		hitMap[i][j][2]++
		res = hitMap[i][j][2]
	case '<':
		hitMap[i][j][3]++
		res = hitMap[i][j][3]
	}
	if res > 1 {
		// fmt.Println("loop")
		return true
	}
	return false
}
