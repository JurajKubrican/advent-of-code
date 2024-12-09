package main

import (
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	id   int
	size int
	pad  int
	// moved bool
}

func a() {
	scanner := getScanner("in.txt")

	scanner.Scan()
	line := scanner.Text()

	parts := strings.Split(line, "")
	files := make([]File, 0)
	for i, p := range parts {
		res, err := strconv.Atoi(p)
		if i%2 == 0 {
			files = append(files, File{size: res, id: i / 2})
		} else {
			files[len(files)-1].pad = res
		}
		if err != nil {
			fmt.Println("Error converting to int", p)
		}
	}

	print(files)

	for i := 0; !isAllOnLeft(files) && i < len(files); {
		j := len(files) - 1
		a := files[i]
		if a.pad == 0 {
			i++
			continue
		}
		b := files[j]
		if a.pad >= b.size {
			pad := a.pad - b.size
			files[i].pad = 0
			b.pad = pad
			files = safeInsert(files, i, b)
			files = files[:len(files)-1]
		} else { // need to split
			movedSpace := a.pad
			files[j].size -= movedSpace
			files[i].pad = 0
			newItem := File{
				id:   b.id,
				size: movedSpace,
				pad:  0,
			}
			fmt.Println("NEW", newItem, j)
			files = safeInsert(files, i, newItem)
		}
	}
	files[len(files)-1].pad = 0
	print(files)

	fmt.Println("DONE", calculateChecksum(files))
}

func print(files []File) {
	for _, f := range files {
		for i := 0; i < f.size; i++ {
			fmt.Print(f.id, " ")

		}
		for i := 0; i < f.pad; i++ {
			fmt.Print(".")
		}

	}

	fmt.Println()
}

func calculateChecksum(files []File) int64 {
	total := int64(0)
	ind := int64(0)
	for _, f := range files {
		for i := 0; i < f.size; i++ {
			total += int64(f.id) * ind
			ind++

		}
	}
	return total

}

func safeInsert(slice []File, index int, element File) []File {
	index++
	if index < 0 || index > len(slice) {
		panic("index out of range")
	}

	// Create a new slice with enough capacity to hold the original slice and the new element
	newSlice := make([]File, len(slice)+1)

	// Copy the elements before the insertion point
	copy(newSlice, slice[:index])

	// Insert the new element
	newSlice[index] = element

	// Copy the elements after the insertion point
	copy(newSlice[index+1:], slice[index:])

	return newSlice
}

func isAllOnLeft(files []File) bool {
	for i, f := range files {
		if f.id >= 0 && f.pad > 0 && i < len(files)-1 {
			return false
		}

	}
	return true
}
