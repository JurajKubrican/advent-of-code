package main

import (
	"fmt"
	"slices"
	"strings"
)

func b() {
	scanner := getScanner("in.txt")

	res := 0
	// scan map

	matrix := make([][]string, 0)
	robot := Point{}
	for scanner.Scan() {
		line := getLine(scanner)
		l := make([]string, 0)
		for _, c := range strings.Split(line, "") {
			if c == "@" {
				l = append(l, "@", ".")
			} else if c == "O" {
				l = append(l, "[", "]")
			} else {
				l = append(l, c, c)
			}
		}
		matrix = append(matrix, l)
		if line == "" {
			break
		}
	}

	// scan movements
	movements := []string{}
	for scanner.Scan() {
		line := getLine(scanner)
		movements = append(movements, strings.Split(line, "")...)
	}

	// find robot
	for i, row := range matrix {
		j := slices.Index(row, "@")
		if j > -1 {
			robot.i = i
			robot.j = j
			matrix[i][j] = "."
		}
	}

	//move robot
	for _, movement := range movements {
		// printMatrix(matrix, robot)
		fmt.Println(movement)

		switch movement {
		case "<":
			canWe := tryMoveLeft2(matrix, robot, true)
			if canWe {
				tryMoveLeft2(matrix, robot, false)
				robot.j--
			}
		case ">":
			canWe := tryMoveRight2(matrix, robot, true)
			if canWe {
				tryMoveRight2(matrix, robot, false)
				robot.j++
			}
		case "^":
			canWe := tryMoveUp2(matrix, robot, true)
			if canWe {
				tryMoveUp2(matrix, robot, false)
				robot.i--
			}
		case "v":
			canWe := tryMoveDown2(matrix, robot, true)
			if canWe {
				tryMoveDown2(matrix, robot, false)
				robot.i++
			}
		}

	}

	printMatrix(matrix, robot)
	fmt.Println(movements)

	for i, row := range matrix {
		for j, c := range row {
			if c == "[" {
				res += i*100 + j
			}
		}
	}

	fmt.Println(res)

}

func tryMoveLeft2(matrix [][]string, point Point, dryRun bool) bool {
	nextPoint := Point{point.i, point.j - 1}
	nextChar := matrix[nextPoint.i][nextPoint.j]

	res := true
	// fmt.Println(" ", point.j, nextChar)

	if nextChar == "#" {
		res = false
	} else if nextChar == "." {
		res = true
	} else if nextChar == "]" || nextChar == "[" {
		res = tryMoveLeft2(matrix, nextPoint, dryRun)
	}

	if !dryRun && res {
		matrix[point.i][point.j], matrix[nextPoint.i][nextPoint.j] = matrix[nextPoint.i][nextPoint.j], matrix[point.i][point.j]
	}

	return res
}

func tryMoveRight2(matrix [][]string, point Point, dryRun bool) bool {
	nextPoint := Point{point.i, point.j + 1}
	nextChar := matrix[nextPoint.i][nextPoint.j]

	res := true
	// fmt.Println(" ", point.j, nextChar)

	if nextChar == "#" {
		res = false
	} else if nextChar == "." {
		res = true
	} else if nextChar == "]" || nextChar == "[" {
		res = tryMoveRight2(matrix, nextPoint, dryRun)
	}

	if !dryRun && res {
		matrix[point.i][point.j], matrix[nextPoint.i][nextPoint.j] = matrix[nextPoint.i][nextPoint.j], matrix[point.i][point.j]
	}

	return res
}

func tryMoveUp2(matrix [][]string, point Point, dryRun bool) bool {
	nextPoint := Point{point.i - 1, point.j}
	nextChar := matrix[nextPoint.i][nextPoint.j]

	res := true
	// fmt.Println(" ", point.j, nextChar)

	if nextChar == "#" {
		res = false
	} else if nextChar == "." {
		res = true
	} else if nextChar == "]" {
		a := tryMoveUp2(matrix, nextPoint, dryRun)
		b := tryMoveUp2(matrix, Point{nextPoint.i, point.j - 1}, dryRun)

		res = a && b
	} else if nextChar == "[" {
		a := tryMoveUp2(matrix, nextPoint, dryRun)
		b := tryMoveUp2(matrix, Point{nextPoint.i, point.j + 1}, dryRun)

		res = a && b
	}

	if !dryRun && res {
		matrix[point.i][point.j], matrix[nextPoint.i][nextPoint.j] = matrix[nextPoint.i][nextPoint.j], matrix[point.i][point.j]
	}

	return res
}

func tryMoveDown2(matrix [][]string, point Point, dryRun bool) bool {
	nextPoint := Point{point.i + 1, point.j}
	nextChar := matrix[nextPoint.i][nextPoint.j]

	res := true
	// fmt.Println(" ", point.j, nextChar)

	if nextChar == "#" {
		res = false
	} else if nextChar == "." {
		res = true
	} else if nextChar == "]" {
		a := tryMoveDown2(matrix, nextPoint, dryRun)
		b := tryMoveDown2(matrix, Point{nextPoint.i, point.j - 1}, dryRun)

		res = a && b
	} else if nextChar == "[" {
		a := tryMoveDown2(matrix, nextPoint, dryRun)
		b := tryMoveDown2(matrix, Point{nextPoint.i, point.j + 1}, dryRun)
		fmt.Println(a, b)
		res = a && b
	}

	if !dryRun && res {
		matrix[point.i][point.j], matrix[nextPoint.i][nextPoint.j] = matrix[nextPoint.i][nextPoint.j], matrix[point.i][point.j]
	}

	return res
}
