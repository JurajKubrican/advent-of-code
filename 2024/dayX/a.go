package main

import "fmt"

func a() {
	scanner := getScanner("in.txt")

	res := 0
	for scanner.Scan() {
		line := getLine(scanner)
		fmt.Println(line)
	}

	fmt.Println(res)

}
