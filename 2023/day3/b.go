package day3

import (
	"fmt"
	"regexp"

	"github.com/JurajKubrican/advent-of-code/2023/files"
)

func Day3B() {

	scanner := files.GetScanner("./day3/test.txt")

	sum := 0

	numberEx, _ := regexp.Compile("[0-9]+")
	starEx, _ := regexp.Compile("[*]")

	numbers := [][]int{}
	stars := [][]int{}

	matrix := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, make([]rune, len(line)))

		numbers = append(numbers, numberEx.FindAllIndex([]byte(line), -1)...)
		stars = append(stars, starEx.FindAllIndex([]byte(line), -1)...)

	}
	fmt.Println("numbers", numbers)
	fmt.Println("stars", stars)

	fmt.Println(sum)

}
