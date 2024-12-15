package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func a() {
	scanner := getScanner("in.txt")

	robots := make([][]int, 0)

	// maxX := 11
	maxX := 101
	// maxY := 7
	maxY := 103
	for scanner.Scan() {
		line := getLine(scanner)
		parts := extractIntegers(line)
		robots = append(robots, parts)
	}
	fmt.Println(robots)

	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		matrix := solve(robots, i, maxX, maxY)
		if ok, _ := findVerticalLineOf5(matrix); ok != -1 {
			printMatrix(matrix)
			continue

		}
		// printMatrix(matrix)
		// time.Sleep(500 * time.Millisecond)
	}

}

func solve(robots [][]int, steps, maxX, maxY int) [][]int {
	matrix := makeMatrix(maxY, maxX)
	for _, robot := range robots {
		x := positiveMod(robot[0]+steps*robot[2], maxX)
		y := positiveMod(robot[1]+steps*robot[3], maxY)
		matrix[y][x]++
	}

	return matrix
}

func makeMatrix(i, j int) [][]int {
	matrix := make([][]int, i)
	for k := range matrix {
		matrix[k] = make([]int, j)
	}
	return matrix
}

func extractIntegers(line string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(line, -1)

	integers := make([]int, len(matches))
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println("Error converting to int:", match)
			continue
		}
		integers[i] = num
	}
	return integers
}

func printMatrix(matrix [][]int) {
	for y, row := range matrix {
		for x := range row {
			cell := matrix[y][x]
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()

	}
}

func positiveMod(a, b int) int {
	mod := a % b
	// fmt.Println(a, "%", b, "=", mod)
	if mod < 0 {
		// fmt.Println(mod, "+", b, "=", mod+b)
		return mod + b
	} else {
		return mod
	}
}

func count(matrix [][]int) int {
	lenx := len(matrix[0]) / 2
	leny := len(matrix) / 2

	fmt.Println(lenx, leny)

	a := getPatch(matrix, 0, 0, lenx, leny)
	printMatrix(a)
	println()
	b := getPatch(matrix, 0, leny+1, lenx, leny)
	printMatrix(b)
	println()

	c := getPatch(matrix, lenx+1, 0, lenx, leny)
	printMatrix(c)

	println()
	d := getPatch(matrix, lenx+1, leny+1, lenx, leny)
	printMatrix(d)

	aa := sumMatrix(a)
	bb := sumMatrix(b)
	cc := sumMatrix(c)
	dd := sumMatrix(d)

	fmt.Println(aa, bb, cc, dd)

	return aa * bb * cc * dd
}

func getPatch(matrix [][]int, x, y, jsize, isize int) [][]int {
	patch := make([][]int, isize)
	for i := 0; i < isize; i++ {
		patch[i] = make([]int, jsize)
		for j := 0; j < jsize; j++ {
			patch[i][j] = matrix[y+i][x+j]
		}
	}
	return patch
}

func sumMatrix(matrix [][]int) int {
	sum := 0
	for _, row := range matrix {
		for _, cell := range row {
			sum += cell
		}
	}
	return sum
}

func findVerticalLineOf5(matrix [][]int) (int, int) {
	for y := 0; y < len(matrix)-4; y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] > 0 && matrix[y+1][x] > 0 && matrix[y+2][x] > 0 && matrix[y+3][x] > 0 && matrix[y+4][x] > 0 {
				return x, y
			}
		}
	}
	return -1, -1
}
