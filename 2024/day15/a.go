package main

import (
	"fmt"
	"slices"
	"strings"
)

type Point struct {
	i int
	j int
}

func a() {
	scanner := getScanner("in.txt")

	res := 0
	// scan map

	matrix := make([][]string, 0)
	robot := Point{}
	for scanner.Scan() {
		line := getLine(scanner)
		matrix = append(matrix, strings.Split(line, ""))
		if line == "" {
			break
		}
	}

	// scan moements
	movements := []string{}
	for scanner.Scan() {
		line := getLine(scanner)
		movements = append(movements, strings.Split(line, "")...)
		if line == "" {
			break
		}
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
			robot = tryMoveLeft(matrix, robot)
		case "^":
			robot = tryMoveUp(matrix, robot)
		case ">":
			robot = tryMoveRight(matrix, robot)
		case "v":
			robot = tryMoveDown(matrix, robot)
		}

	}

	printMatrix(matrix, robot)
	fmt.Println(movements)

	for i, row := range matrix {
		for j, c := range row {
			if c == "O" {
				res += i*100 + j
			}
		}
	}

	fmt.Println(res)

}

func printMatrix(matrix [][]string, robot Point) {
	for i, row := range matrix {
		for j, col := range row {
			if robot.i == i && robot.j == j {
				fmt.Print("@")
				continue
			}
			fmt.Print(col)
		}
		fmt.Println()
	}
}

func tryMoveLeft(matrix [][]string, robot Point) Point {
	carrying := Point{-1, -1}
	i := robot.i
	for j := robot.j - 1; ; j-- {
		moved, done, c := iterateMovement(matrix, i, j, carrying)
		carrying = c
		fmt.Println(moved, done, c)
		if moved {
			robot.j--
		}
		if done {
			return robot
		}

	}
}
func tryMoveRight(matrix [][]string, robot Point) Point {
	carrying := Point{-1, -1}
	i := robot.i
	for j := robot.j + 1; ; j++ {
		moved, done, c := iterateMovement(matrix, i, j, carrying)
		carrying = c
		if moved {
			robot.j++
		}
		if done {
			return robot
		}

	}
}
func tryMoveUp(matrix [][]string, robot Point) Point {
	carrying := Point{-1, -1}
	j := robot.j
	for i := robot.i - 1; ; i-- {
		moved, done, c := iterateMovement(matrix, i, j, carrying)
		carrying = c
		if moved {
			robot.i--
		}
		if done {
			return robot
		}
	}
}
func tryMoveDown(matrix [][]string, robot Point) Point {
	carrying := Point{-1, -1}
	j := robot.j
	for i := robot.i + 1; ; i++ {
		moved, done, c := iterateMovement(matrix, i, j, carrying)
		carrying = c
		if moved {
			robot.i++
		}
		if done {
			return robot
		}
	}
}

func iterateMovement(matrix [][]string, i, j int, carrying Point) (bool, bool, Point) {
	if matrix[i][j] == "#" {
		return false, true, carrying
	}

	if matrix[i][j] == "O" && carrying.i == -1 {
		carrying = Point{i, j}
	}

	if matrix[i][j] == "." {
		if carrying.i > -1 {
			matrix[carrying.i][carrying.j] = "."
			matrix[i][j] = "O"
		}
		// move
		return true, true, carrying
	}

	return false, false, carrying
}
