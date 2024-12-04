package main

import (
	"fmt"
	"strconv"
)

func a() {
	scanner := getScanner("in.txt")

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
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
