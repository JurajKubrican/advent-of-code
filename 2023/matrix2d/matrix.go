package matrix2d

import (
	"fmt"
	"strings"
)

func Get9Patch(matrix [][]string, iStart, jStart int, fallback string) [][]string {

	patch := [][]string{
		{fallback, fallback, fallback},
		{fallback, fallback, fallback},
		{fallback, fallback, fallback},
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			patch[i+1][j+1] = getOrDefault(matrix, iStart+i, jStart+j, fallback)
		}
	}

	return patch
}

func isIndexInBounds(matrix [][]string, i, j int) bool {
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

func getOrDefault(matrix [][]string, i, j int, fallback string) string {
	if isIndexInBounds(matrix, i, j) {
		return matrix[i][j]
	}
	return fallback
}

func Contains(matrix [][]string, contains string) bool {

	for _, line := range matrix {
		for _, char := range line {
			if strings.Contains(contains, char) {
				return true
			}
		}
	}

	return false
}

func AdjacentIndicesDiagonal(matrix [][]string, i, j int) [][]int {
	indices := [][]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},

		{i, j - 1},
		{i, j + 1},

		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}

	res := [][]int{}
	for _, val := range indices {
		if val[0] > 0 && val[1] > 0 && val[0] <= len(matrix) && val[1] <= len(matrix) {
			res = append(res, val)
		}
	}

	return res
}

func PrintMatrix(matrix [][]string) {
	for _, line := range matrix {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Println()
	}
}
