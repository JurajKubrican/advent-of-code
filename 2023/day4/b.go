package day4

import (
	"fmt"
	"strings"

	"github.com/JurajKubrican/advent-of-code/2023/files"
)

func Day4B() {

	scanner := files.GetScanner("./day4/in.txt")

	sum := 0
	copies := map[int]int{}

	i := 0
	for scanner.Scan() {
		copies[i] += 1
		mul := copies[i]
		i++

		line := scanner.Text()
		parts := strings.Split(line, ":")
		parts = strings.Split(parts[1], " | ")
		wining := strings.Fields(parts[0])
		guesses := strings.Fields(parts[1])

		intersection := intersect(wining, guesses)

		count := len(intersection)
		for j := range count {
			copies[i+j] += mul
		}

		fmt.Println(i)
	}
	for _, i := range copies {
		sum += i
	}

	fmt.Println(copies)
	fmt.Println(sum)

}
