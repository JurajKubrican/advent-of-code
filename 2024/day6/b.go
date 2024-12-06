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

}

func traverseIsLoop(matrix [][]rune, guard []int, guardDir rune) bool {
	hitMap := make([][][]bool, len(matrix))

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
			continue
		}
		if next == '.' || next == 'X' || next == '^' {
			guard[0], guard[1] = proceed(guard[0], guard[1], guardDir)
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

func countHitFromDirection(hitMap [][][]bool, i, j int, dir rune) bool {
	if hitMap[i] == nil {
		hitMap[i] = make([][]bool, len(hitMap))
	}
	if hitMap[i][j] == nil {
		hitMap[i][j] = []bool{false, false, false, false}
	}
	switch dir {
	case '^':
		if hitMap[i][j][0] {
			return true
		}
		hitMap[i][j][0] = true
	case '>':
		if hitMap[i][j][1] {
			return true
		}
		hitMap[i][j][1] = true
	case 'v':
		if hitMap[i][j][2] {
			return true
		}
		hitMap[i][j][2] = true
	case '<':
		if hitMap[i][j][3] {
			return true
		}
		hitMap[i][j][3] = true
	}

	return false
}
