package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	i int
	j int
}

func a() {
	scanner := getScanner("in.txt")

	total := 0
	matrix := make([][]int, 0)

	for scanner.Scan() {
		line := getLine(scanner)
		row := parseIntSlice(line, "")
		matrix = append(matrix, row)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				continue
			}
			fmt.Println("START", i, j, matrix[i][j])
			results := traverse(matrix, i, j)
			if len(results) > 0 {
				total += len(results)
			}
			fmt.Println("RESULTS", results)

		}
	}

	printMatrix(matrix)
	fmt.Println("DONE", total)

}

func traverse(matrix [][]int, i, j int) map[Point]int {
	if matrix[i][j] == 9 {
		return map[Point]int{{i, j}: 0}
	}
	if getOrDefault(matrix, i, j, -1) == -1 {
		return map[Point]int{}
	}
	fmt.Println(i, j, matrix[i][j])

	results := map[Point]int{}
	fmt.Println(matrix[i][j], getOrDefault(matrix, i+1, j, -99)+1)

	if matrix[i][j]+1 == getOrDefault(matrix, i+1, j, -99) {
		mergeMaps(results, traverse(matrix, i+1, j))
	}
	if matrix[i][j]+1 == getOrDefault(matrix, i, j+1, -99) {
		mergeMaps(results, traverse(matrix, i, j+1))
	}
	if matrix[i][j]+1 == getOrDefault(matrix, i-1, j, -99) {
		mergeMaps(results, traverse(matrix, i-1, j))
	}
	if matrix[i][j]+1 == getOrDefault(matrix, i, j-1, -99) {
		mergeMaps(results, traverse(matrix, i, j-1))
	}
	return results
}

func parseIntSlice(s string, separator string) []int {
	parts := strings.Split(s, separator)
	result := make([]int, len(parts))
	for i, p := range parts {
		res, err := strconv.Atoi(p)

		if err != nil {
			result[i] = -99
			fmt.Println("Error converting to int", p)
		} else {
			result[i] = res
		}
	}
	return result
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func getPatchCopy(iStart, jStart int, matrix [][]int) [][]int {
	patch := make([][]int, 3)
	for i := -1; i <= 1; i++ {
		patch[i+1] = make([]int, 3)
		for j := -1; j <= 1; j++ {
			patch[i+1][j+1] = getOrDefault(matrix, iStart+i, jStart+j, -1)
		}
	}
	return patch
}

func isIndexInBounds(i, j int, matrix [][]int) bool {
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

func getOrDefault(matrix [][]int, i, j int, fallback int) int {
	if isIndexInBounds(i, j, matrix) {
		return matrix[i][j]
	}
	return fallback
}

func mergeMaps(dest, src map[Point]int) {
	for k, v := range src {
		dest[k] = v
	}
}
