package main

import (
	"fmt"
)

func b() {
	scanner := getScanner("in.txt")

	total := 0

	matrix := make([][]rune, 0)
	i := 0
	for scanner.Scan() {
		line := getLine(scanner)
		matrix = append(matrix, make([]rune, len(line)))
		for j, c := range line {
			matrix[i][j] = c
		}
		i++
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 'A' {
				continue
			}
			patch := getPatch(i, j, matrix)
			if patch[0][0] == 'M' && patch[0][2] == 'M' && patch[2][0] == 'S' && patch[2][2] == 'S' {
				total++
				fmt.Println("FOUND1", i, j)
				printRuneMatrix(patch)

			}
			if patch[0][0] == 'S' && patch[0][2] == 'S' && patch[2][0] == 'M' && patch[2][2] == 'M' {
				total++
				fmt.Println("FOUND2", i, j)
				printRuneMatrix(patch)
			}
			if patch[0][0] == 'M' && patch[0][2] == 'S' && patch[2][0] == 'M' && patch[2][2] == 'S' {
				total++
				fmt.Println("FOUND3", i, j)
				printRuneMatrix(patch)
			}
			if patch[0][0] == 'S' && patch[0][2] == 'M' && patch[2][0] == 'S' && patch[2][2] == 'M' {
				total++
				fmt.Println("FOUND4", i, j)
				printRuneMatrix(patch)
			}

			// printRuneMatrix(patch)
		}
	}

	fmt.Println("DONE", total)

}

func getPatch(iStart, jStart int, matrix [][]rune) [][]rune {
	patch := make([][]rune, 3)
	for i := -1; i <= 1; i++ {
		patch[i+1] = make([]rune, 3)
		for j := -1; j <= 1; j++ {
			patch[i+1][j+1] = getOrDefault(iStart+i, jStart+j, matrix, '.')
		}
	}
	return patch
}
