package main

import (
	"fmt"
	"math"
	"slices"
)

func a() {
	scanner := getScanner("in.txt")

	listA := []int{}
	listB := []int{}

	for scanner.Scan() {
		line := getLineOfInts(scanner)
		fmt.Println(line)
		listA = append(listA, line[0])
		listB = append(listB, line[1])
	}
	slices.Sort(listA)
	slices.Sort(listB)

	sum := 0
	for i := 0; i < len(listA); i++ {
		sum += int(math.Abs(float64(listA[i] - listB[i])))
	}

	fmt.Println(listA)
	fmt.Println(listB)
	fmt.Println(sum)

}
