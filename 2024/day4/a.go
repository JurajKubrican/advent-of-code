package main

import (
	"fmt"
)

func a() {
	scanner := getScanner("in.txt")

	WORD_SIZE := 4

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
			words := getStar(i, j, WORD_SIZE, matrix)
			for _, word := range words {
				if string(word) == "XMAS" {
					total++
				}
			}
		}
	}

	fmt.Println("DONE", total)

}

func getStar(iStart, jStart, size int, matrix [][]rune) [][]rune {
	words := make([][]rune, 8)
	for i := 0; i < size; i++ {
		words[0] = append(words[0], getOrDefault(iStart, jStart+i, matrix, '.'))
		words[1] = append(words[1], getOrDefault(iStart+i, jStart, matrix, '.'))
		words[2] = append(words[2], getOrDefault(iStart, jStart-i, matrix, '.'))
		words[3] = append(words[3], getOrDefault(iStart-i, jStart, matrix, '.'))

		words[4] = append(words[4], getOrDefault(iStart+i, jStart+i, matrix, '.'))
		words[5] = append(words[5], getOrDefault(iStart+i, jStart-i, matrix, '.'))
		words[6] = append(words[6], getOrDefault(iStart-i, jStart-i, matrix, '.'))
		words[7] = append(words[7], getOrDefault(iStart-i, jStart+i, matrix, '.'))
	}
	return words
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

func printRuneMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}
