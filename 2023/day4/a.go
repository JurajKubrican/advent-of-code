package day4

import (
	"fmt"
	"math"
	"strings"

	"github.com/JurajKubrican/advent-of-code/2023/files"
)

func Day4A() {

	scanner := files.GetScanner("./day4/in.txt")

	sum := 0.0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		parts = strings.Split(parts[1], " | ")
		wining := strings.Fields(parts[0])
		guesses := strings.Fields(parts[1])

		intersection := intersect(wining, guesses)

		count := float64(len(intersection))

		if count > 0 {
			sum += math.Pow(2, count-1)
		}

	}

	fmt.Println(sum)

}

func intersect(a, b []string) []string {
	hashmap := make(map[string]struct{})
	for _, i := range a {
		hashmap[i] = struct{}{}
	}

	res := make([]string, 0)
	for _, i := range b {
		if _, ok := hashmap[i]; ok {
			res = append(res, i)
		}
	}
	return res
}
