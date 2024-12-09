package main

import (
	"fmt"
	"strconv"
	"strings"
)

func b() {
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

	fmt.Println(files)
	// print(files)
	// fmt.Println("BEFORE", calculateChecksum2(files))
	for i := len(files) - 1; i >= 0; i-- {
		fmt.Print(files[i].size, ":")

		// fmt.Println("FINDING PLACE FOR: ", files[i])
		for j := 0; j < i; j++ {
			// fmt.Print(files[j].pad, " ")
			if files[j].pad >= files[i].size {

				if i == j+1 {
					// fmt.Println("THIS FUCKER")
					s := files[i].size
					files[i-1].pad -= s
					files[i].pad += s
					i++
					break
				}

				// fmt.Println("Moving", files[i].id, "from", i, "to", j+1)
				newItemPad := files[j].pad - files[i].size

				padAfterLeaving := files[i-1].pad + files[i].size + files[i].pad
				files[i-1].pad = padAfterLeaving
				files[i].pad = newItemPad
				files[j].pad = 0

				files = SafeMove(files, i, j+1)
				// print(files)
				i++
				break
			}
		}
		// print(files)
		// fmt.Println(files)

	}

	fmt.Println(len(files))

	// print(files)

	fmt.Println("AFTER", calculateChecksum2(files))
}

// SafeMove moves an element from one index to another in a slice without modifying the original slice.
func SafeMove(slice []File, fromIndex, toIndex int) []File {

	// Create a new slice with the same length as the original slice
	newSlice := make([]File, len(slice))

	// Copy the elements before the fromIndex
	copy(newSlice, slice[:fromIndex])

	// Copy the elements after the fromIndex
	copy(newSlice[fromIndex:], slice[fromIndex+1:])

	// Insert the element at the toIndex
	if toIndex < fromIndex {
		copy(newSlice[toIndex+1:], newSlice[toIndex:])
		newSlice[toIndex] = slice[fromIndex]
	} else {
		copy(newSlice[toIndex:], newSlice[toIndex+1:])
		newSlice[toIndex] = slice[fromIndex]
	}

	return newSlice
}

func calculateChecksum2(files []File) int64 {
	total := int64(0)
	ind := int64(0)

	for _, f := range files {
		for i := 0; i < f.size; i++ {
			total += int64(f.id) * ind
			ind++

		}
		for i := 0; i < f.pad; i++ {
			ind++
		}
	}

	fmt.Println(ind)
	return total

}
