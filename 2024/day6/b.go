package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

const (
	Up    = 1 << iota // 1 << 0 which is 0001
	Down              // 1 << 1 which is 0010
	Left              // 1 << 2 which is 0100
	Right             // 1 << 3 which is 1000
)

func b() {
	inputFile := "in.txt"
	if len(os.Args) == 2 {
		inputFile = os.Args[1]
	}

	scanner := getScanner(inputFile)
	total := 0

	// LOAD MATRIX
	matrix := make([][]byte, 0)
	guard := make([]int, 2)

	for scanner.Scan() {
		line := getLine(scanner)
		parts := []byte(line)

		index := slices.Index(parts, '^')
		if index != -1 {
			guard[0] = len(matrix)
			guard[1] = index
		}
		matrix = append(matrix, parts)
	}

	prevChar := byte(' ')
	isFirst := true
	for i := range matrix {
		for j := range matrix[i] {
			if !isFirst && (matrix[i][j] == '#' || matrix[i][j] == '.') {
				continue
			}

			tempGuard := make([]int, 2)
			copy(tempGuard, guard)
			tempGuardDir := byte('^')
			prevChar = matrix[i][j]
			matrix[i][j] = '#'

			result := traverseIsLoop(matrix, tempGuard, tempGuardDir, isFirst)
			if result {
				total++
			}
			matrix[i][j] = prevChar
			isFirst = false

		}
	}
	// printMatrix(matrix)

	fmt.Println(total)

}

func traverseIsLoop(matrix [][]byte, guard []int, guardDir byte, mark bool) bool {
	// hitMap := make([][][]bool, len(matrix))
	hitMap := map[int]byte{}

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
			if mark {
				matrix[guard[0]][guard[1]] = 'X'
			}
			continue
		}

	}
	return false
}

func countHitFromDirection(hitMap map[int]byte, i, j int, dir byte) bool {
	ind := i*1000 + j

	var direction byte
	switch dir {
	case '^':
		direction = Up
	case 'v':
		direction = Down
	case '<':
		direction = Left
	case '>':
		direction = Right
	}

	if hitMap[ind]&direction != 0 {
		return true
	}

	hitMap[ind] |= direction
	return false
}
