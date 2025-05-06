package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var A = 0
var B = 0
var C = 0

var index = 0

var program = make([]int, 0)

var output = make([]int, 0)

func combo(i int) int {
	if i == 4 {
		return A
	}
	if i == 5 {
		return B
	}
	if i == 6 {
		return C
	}
	return i
}

func adv() {
	arg := combo(program[index+1])
	a := A
	b := int(math.Pow(2, float64(arg)))
	res := a / b
	A = res
	index += 2
}
func bxl() {
	arg := program[index+1]
	a := B
	b := arg
	res := a ^ b
	B = res
	index += 2
}
func bst() {
	arg := combo(program[index+1])
	a := arg
	res := a % 8
	B = res
	index += 2
}
func jnz() {
	arg := program[index+1]
	if A == 0 {
		index += 2
	} else {
		index = arg
	}

}
func bxc() {
	res := B ^ C
	B = res
	index += 2
}
func out() {
	arg := combo(program[index+1])
	res := arg % 8
	output = append(output, res)
	index += 2
}
func bdv() {
	arg := combo(program[index+1])
	a := A
	b := int(math.Pow(2, float64(arg)))
	res := a / b
	B = res
	index += 2
}
func cdv() {
	arg := combo(program[index+1])
	a := A
	b := int(math.Pow(2, float64(arg)))
	res := a / b
	C = res
	index += 2
}

var instructionMap = []func(){
	adv,
	bxl,
	bst,
	jnz,
	bxc,
	out,
	bdv,
	cdv,
}

func fna() {
	scanner := getScanner("in.txt")

	res := 0

	scanner.Scan()
	line := getLine(scanner)
	line = strings.Split(line, "Register A: ")[1]
	A, _ = strconv.Atoi(line)

	scanner.Scan()
	line = getLine(scanner)
	line = strings.Split(line, "Register B: ")[1]
	B, _ = strconv.Atoi(line)

	scanner.Scan()
	line = getLine(scanner)
	line = strings.Split(line, "Register C: ")[1]
	C, _ = strconv.Atoi(line)

	scanner.Scan()
	scanner.Scan()
	line = getLine(scanner)
	sprogram := strings.Split(strings.Split(line, "Program: ")[1], ",")
	for _, v := range sprogram {
		i, _ := strconv.Atoi(v)
		program = append(program, i)
	}

	fmt.Println(A, B, C, program)

	fmt.Println(res)

	runProgramSimple()

	fmt.Println(A, B, C, output)

	strOutput := []string{}
	for _, v := range output {
		strOutput = append(strOutput, strconv.Itoa(v))
	}

	fmt.Println(strings.Join(strOutput, ","))

}

func runProgram() {
	// for index < len(program) {
	// 	instructionMap[program[index]]()
	// }
}

func runProgramSimple() {
	for {
		B := (A % 8) % 3
		C := A / 32
		B = B ^ 5
		A := A / 8
		B = B ^ C
		output = append(output, B%8)
		if A == 0 {
			break
		}
	}
}
