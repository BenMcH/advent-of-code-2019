package intcode

import (
	"advent-of-code-2019/utils"
	"fmt"
	"strings"
)

func ExecuteIntcodeFromString(intcode string) ([]int, []int) {
	program, _ := utils.StringsToInts(strings.Split(intcode, ","))
	return ExecuteIntcode(program)
}

func ExecuteIntcode(intcode []int) (program []int, output []int) {
	program = make([]int, len(intcode))
	copy(program, intcode)

	output = []int{}

	pos := 0
	for pos < len(program) {
		switch program[pos] {
		case 1:
			args := program[pos : pos+4]
			a, b, nPos := program[args[1]], program[args[2]], args[3]
			program[nPos] = a + b
			pos += 4
		case 2:
			args := program[pos : pos+4]
			a, b, nPos := program[args[1]], program[args[2]], args[3]
			program[nPos] = a * b
			pos += 4
		case 3:
			input := 1
			param := program[pos+1]
			pos += 2

			program[param] = input
		case 4:
			param := program[pos+1]
			value := program[param]
			pos += 2

			output = append(output, value)
		case 99:
			pos = len(program)
		default:
			panic(fmt.Sprintf("Unknown opcode %d", program[pos]))
		}
	}

	return program, output
}
