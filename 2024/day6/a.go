package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func a() {
	scanner := getScanner("in.txt")

	total := 0

	matrix := make([][]rune, 0)
	guard := make([]int, 2)
	guardDir := '^'

	for scanner.Scan() {
		line := getLine(scanner)
		parts := []rune(line)
		fmt.Println(parts)
		index := slices.Index(parts, '^')
		if index != -1 {
			guard[0] = len(matrix)
			guard[1] = index
		}
		matrix = append(matrix, parts)
	}

	for true {
		next := getNeighbor(matrix, guard[0], guard[1], guardDir)
		if next == '~' {
			break //done
		}
		if next == '#' {
			guardDir = nextDirection(guardDir)
			fmt.Println("TURN", guardDir)
			continue
		}
		if next == '.' || next == 'X' || next == '^' {

			guard[0], guard[1] = proceed(guard[0], guard[1], guardDir)
			fmt.Println("PROCEED", guardDir)
			matrix[guard[0]][guard[1]] = 'X'
			continue
		}

	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 'X' || matrix[i][j] == '^' {
				total++
			}
		}
	}

	fmt.Println("DONE", total, guard)
	printMatrix(matrix)

}

func getPatchCopy(iStart, jStart int, matrix [][]rune) [][]rune {
	patch := make([][]rune, 3)
	for i := -1; i <= 1; i++ {
		patch[i+1] = make([]rune, 3)
		for j := -1; j <= 1; j++ {
			patch[i+1][j+1] = getOrDefault(iStart+i, jStart+j, matrix, '.')
		}
	}
	return patch
}

func isIndexInBounds(i, j int, matrix [][]rune) bool {
	if i < 0 || j < 0 {
		return false
	}
	maxI := len(matrix)
	if i >= maxI {
		return false
	}
	maxJ := len(matrix[i])
	return i >= 0 && j >= 0 && i < maxI && j < maxJ
}

func getOrDefault(i, j int, matrix [][]rune, fallback rune) rune {
	if isIndexInBounds(i, j, matrix) {
		return matrix[i][j]
	}
	return fallback
}

func printMatrix(matrix [][]rune) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

}

func getNeighbor(matrix [][]rune, i, j int, dir rune) rune {
	if dir == '^' {
		return getOrDefault(i-1, j, matrix, '~')
	}
	if dir == 'v' {
		return getOrDefault(i+1, j, matrix, '~')
	}
	if dir == '<' {
		return getOrDefault(i, j-1, matrix, '~')
	}
	if dir == '>' {
		return getOrDefault(i, j+1, matrix, '~')
	}

	fmt.Println("UNKNOWN DIRECTION", dir)
	return '~'
}

func nextDirection(dir rune) rune {
	if dir == '^' {
		return '>'
	}
	if dir == '>' {
		return 'v'
	}
	if dir == 'v' {
		return '<'
	}
	if dir == '<' {
		return '^'
	}
	fmt.Println("UNKNOWN DIRECTION", dir)
	return '~'
}

func proceed(i, j int, dir rune) (int, int) {
	if dir == '^' {
		return i - 1, j
	}
	if dir == 'v' {
		return i + 1, j
	}
	if dir == '<' {
		return i, j - 1
	}
	if dir == '>' {
		return i, j + 1
	}

	fmt.Println("UNKNOWN DIRECTION", dir)
	return i, j
}
