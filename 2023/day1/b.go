package main

import (
	"fmt"
	"strconv"
	"strings"
)

var wordToDigit = map[string]string{
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func b() {
	scanner := getScanner("in.txt")

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		first := ""
		last := ""

		// first
		for len(line) > 0 {
			found := false
			for word, digit := range wordToDigit {
				found = strings.HasPrefix(line, word)

				if found {
					first = digit
					break
				}
			}
			if !found {
				line = line[1:]
				fmt.Println(line)
			} else {
				break
			}
		}

		for len(line) > 0 {
			found := false
			for word, digit := range wordToDigit {
				found = strings.HasSuffix(line, word)

				if found {
					last = digit
					break
				}
			}
			if !found {
				line = line[:len(line)-1]
			} else {
				break
			}
		}

		fmt.Println(first, last)
		res, err := strconv.Atoi(first + last)
		if err != nil {
			fmt.Println(err)
		}
		sum += res

	}

	fmt.Println(sum)
}
