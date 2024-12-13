package solver

import (
	"regexp"
	"strconv"
	"strings"
)

type BaseDayThreeSolver struct {
	instructions []string
}

func (solver *BaseDayThreeSolver) ReadData(data []byte) error {
	solver.instructions = strings.Split(strings.TrimSpace(string(data)), "\n")

	return nil
}

type DayThreeSolverP1 struct {
	BaseDayThreeSolver
}

func (solver *DayThreeSolverP1) Solve() []byte {
	reg, _ := regexp.Compile(`mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)`)
	result := 0

	for _, instruction := range solver.instructions {
		matches := reg.FindAllString(instruction, -1)

		for _, match := range matches {
			match = match[4:]
			match = match[:len(match)-1]

			valueStrings := strings.Split(match, ",")
			lhs, _ := strconv.Atoi(valueStrings[0])
			rhs, _ := strconv.Atoi(valueStrings[1])

			result += lhs * rhs
		}
	}

	return []byte(strconv.Itoa(result))
}

type DayThreeSolverP2 struct {
	BaseDayThreeSolver
}

func (solver *DayThreeSolverP2) Solve() []byte {
	reg, _ := regexp.Compile(`(mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\))|(do(n't)?\(\))`)
	result := 0

	isActive := true
	for _, instruction := range solver.instructions {
		matches := reg.FindAllString(instruction, -1)

		for _, match := range matches {
			if match == "do()" {
				isActive = true
				continue
			} else if match == "don't()" {
				isActive = false
				continue
			}

			if !isActive {
				continue
			}

			match = match[4:]
			match = match[:len(match)-1]

			valueStrings := strings.Split(match, ",")
			lhs, _ := strconv.Atoi(valueStrings[0])
			rhs, _ := strconv.Atoi(valueStrings[1])

			result += lhs * rhs
		}
	}

	return []byte(strconv.Itoa(result))
}
