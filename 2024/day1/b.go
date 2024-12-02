package main

import (
	"fmt"
	"slices"
)

func b() {
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
		a := listA[i]
		count := getCountOf(listB, a)
		sum += count * a
	}

	fmt.Println(listA)
	fmt.Println(listB)
	fmt.Println(sum)

}

func getCountOf(list []int, num int) int {
	count := 0
	for _, i := range list {
		if i == num {
			count++
		}
	}
	return count
}
