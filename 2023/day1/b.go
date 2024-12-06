package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func b() {
	scanner := getScanner("in.txt")

	sum := 0

	// Define a map of words to digits

	// Create a regular expression pattern to match any of the words
	pattern := regexp.MustCompile(strings.Join(keys(wordToDigit), "|"))

	for scanner.Scan() {
		line := scanner.Text()

		// Replace words with corresponding digits using the regular expression
		line = pattern.ReplaceAllStringFunc(line, func(matched string) string {
			return wordToDigit[matched]
		})

		fmt.Println(line)
		a := 0
		for _, c := range line {
			num, err := strconv.Atoi(string(c))
			if err == nil {
				a = num
				fmt.Println(num)
				break
			}
		}
		b := 0
		for i := range line {
			c := line[len(line)-i-1]
			num, err := strconv.Atoi(string(c))
			if err == nil {
				b = num
				fmt.Println(num)
				break
			}
		}
		res, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
		sum += res

	}

	fmt.Println(sum)

}




func getNum(in string) int {

	wordToDigit := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":  6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	
	for key, value := range (wordToDigit)
 if string.HasPrefix()
}
