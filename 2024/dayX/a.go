package main

func a() {
	scanner := getScanner("in.txt")

	matrix := make([][]Point, 0)
	i := 0
	for scanner.Scan() {
		line := getLine(scanner)
		matrix = append(matrix, make([]Point, len(line)))
		for j, c := range line {
			matrix[i][j] = Point{c: c, perimeter: 0}
		}
		i++
	}

	res := 0

	for i, row := range matrix {
		for j, c := range row {
			a, p := compute(matrix, i, j, c)

			res += a * p

		}
	}

	printM(matrix)
	println()
	printI(matrix)
	println("RES", res)

}

func compute(matrix [][]Point, i, j int, g Point) (int, int) {
	if matrix[i][j].c != g.c {
		return 0, 0
	}
	if matrix[i][j].visited == true {
		return 0, 0
	}

	area := 1
	perimeter := getPerimeter(matrix, i, j)
	matrix[i][j].perimeter = perimeter
	matrix[i][j].visited = true

	if isIndexInBounds(matrix, i-1, j) {
		a, b := compute(matrix, i-1, j, g)
		area += a
		perimeter += b
	}
	if isIndexInBounds(matrix, i+1, j) {
		a, b := compute(matrix, i+1, j, g)
		area += a
		perimeter += b
	}
	if isIndexInBounds(matrix, i, j-1) {
		a, b := compute(matrix, i, j-1, g)
		area += a
		perimeter += b
	}
	if isIndexInBounds(matrix, i, j+1) {
		a, b := compute(matrix, i, j+1, g)
		area += a
		perimeter += b
	}
	return area, perimeter

}

func getPerimeter(matrix [][]Point, i, j int) int {
	patch := getPatchCopy(matrix, i, j, Point{c: ' '})
	c := patch[1][1].c
	res := 0
	if patch[0][1].c != c {
		res++
	}
	if patch[1][0].c != c {
		res++
	}
	if patch[1][2].c != c {
		res++
	}
	if patch[2][1].c != c {
		res++
	}
	return res
}

func printM(matrix [][]Point) {
	for _, row := range matrix {
		for _, c := range row {
			print(string(c.c))
		}
		println()
	}
}

func printP(matrix [][]Point) {
	for _, row := range matrix {
		for _, c := range row {
			print(c.perimeter)
		}
		println()
	}
}

func getPatchCopy(matrix [][]Point, iStart, jStart int, fallback Point) [][]Point {
	patch := make([][]Point, 3)
	for i := -1; i <= 1; i++ {
		patch[i+1] = make([]Point, 3)
		for j := -1; j <= 1; j++ {
			patch[i+1][j+1] = getOrDefault(matrix, iStart+i, jStart+j, fallback)
		}
	}
	return patch
}

func isIndexInBounds(matrix [][]Point, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	maxI := len(matrix)
	if i >= maxI {
		return false
	}
	maxJ := len(matrix[i])
	return i >= 0 && j >= 0 && i < maxI && j < maxJ
}

func getOrDefault(matrix [][]Point, i, j int, fallback Point) Point {
	if isIndexInBounds(matrix, i, j) {
		return matrix[i][j]
	}
	return fallback
}
