package main

import (
	"fmt"
)

func b() {
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
			results := traverse2(matrix, i, j)
			if len(results) > 0 {
				total += countResult(results)
			}
			fmt.Println("RESULTS", results)

		}
	}

	printMatrix(matrix)
	fmt.Println("DONE", total)
}

func traverse2(matrix [][]int, i, j int) map[Point]int {
	if matrix[i][j] == 9 {
		return map[Point]int{{i, j}: 1}
	}
	if getOrDefault(matrix, i, j, -1) == -1 {
		return map[Point]int{}
	}
	fmt.Println(i, j, matrix[i][j])

	results := map[Point]int{}
	fmt.Println(matrix[i][j], getOrDefault(matrix, i+1, j, -99)+1)

	if matrix[i][j]+1 == getOrDefault(matrix, i+1, j, -99) {
		mergeMaps2(results, traverse2(matrix, i+1, j))
		fmt.Println("1", results)
	}
	if matrix[i][j]+1 == getOrDefault(matrix, i, j+1, -99) {
		mergeMaps2(results, traverse2(matrix, i, j+1))
		fmt.Println("2", results)
	}
	if matrix[i][j]+1 == getOrDefault(matrix, i-1, j, -99) {
		mergeMaps2(results, traverse2(matrix, i-1, j))
		fmt.Println("3", results)
	}
	if matrix[i][j]+1 == getOrDefault(matrix, i, j-1, -99) {
		mergeMaps2(results, traverse2(matrix, i, j-1))
		fmt.Println("4", results)
	}
	return results
}

func mergeMaps2(dest, src map[Point]int) {
	for k, v := range src {
		if _, ok := dest[k]; !ok {
			dest[k] = v
		} else {
			dest[k] += v
		}
	}
}

func countResult(results map[Point]int) int {
	total := 0
	for _, v := range results {
		total += v
	}
	return total
}
