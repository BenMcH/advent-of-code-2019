package day_02

import (
	"advent-of-code-2019/utils"
	"fmt"
	"strings"
)

func ExecuteIntcodeFromString(intcode string) []int {
	program, _ := utils.StringsToInts(strings.Split(intcode, ","))
	return ExecuteIntcode(program)
}

func ExecuteIntcode(intcode []int) []int {
	program := make([]int, len(intcode))
	copy(program, intcode)

	for pos := 0; pos < len(program); pos += 4 {
		if pos+3 >= len(program) {
			break
		}

		args := program[pos : pos+4]

		switch args[0] {
		case 1:
			a, b, nPos := program[args[1]], program[args[2]], args[3]
			program[nPos] = a + b
		case 2:
			a, b, nPos := program[args[1]], program[args[2]], args[3]
			program[nPos] = a * b
		case 99:
			pos = len(program)
		default:
			panic(fmt.Sprintf("Unknown opcode %d", args[0]))
		}
	}

	return program
}

func PartOne(intcode string) (int, error) {
	args := strings.Split(intcode, ",")
	ints, _ := utils.StringsToInts(args)

	ints[1] = 12
	ints[2] = 2
	ints = ExecuteIntcode(ints)

	return ints[0], nil
}
