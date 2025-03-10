package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JurajKubrican/advent-of-code/2023/files"
	"github.com/JurajKubrican/advent-of-code/2023/matrix2d"
)

const (
	BLANK = iota
	NUMBER
	VALID_NUMBER
)

func Day3A() {

	scanner := files.GetScanner("./day3/in.txt")

	sum := 0
	matrix := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	state := BLANK
	number := ""

	for i, line := range matrix {
		for j, char := range line {
			if !strings.Contains("0987654321", char) {
				if state == VALID_NUMBER {
					n, _ := strconv.Atoi(number)
					sum += n

				}
				number = ""
				state = BLANK
				continue
			}

			number += char
			patch := matrix2d.Get9Patch(matrix, i, j, ".")
			if matrix2d.Contains(patch, "*+/-=@$%#&") {
				state = VALID_NUMBER
			}

		}

		if state == VALID_NUMBER {
			n, _ := strconv.Atoi(number)
			fmt.Println("+ ", number)
			sum += n
			number = ""
			state = BLANK
		}
	}

	fmt.Println(sum)

}
