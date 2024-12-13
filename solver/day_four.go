package solver

import (
	"strconv"
	"strings"
)

type BaseDayFourSolver struct {
	matrix []string
}

func (solver *BaseDayFourSolver) ReadData(data []byte) error {
	solver.matrix = strings.Split(strings.TrimSpace(string(data)), "\n")

	return nil
}

type DayFourSolverP1 struct {
	BaseDayFourSolver
}

func findIndexes(matrix []string, target rune) []Coord2D {
	var indexes []Coord2D
	for x, rowData := range matrix {
		for y, char := range rowData {
			if char == target {
				indexes = append(indexes, Coord2D{x, y})
			}
		}
	}
	return indexes
}

func isOutOfBounds(position Coord2D, bottomBound int, rightBound int) bool {
	return position.x < 0 ||
		position.x >= bottomBound ||
		position.y < 0 ||
		position.y >= rightBound
}

func countXmas(matrix []string, startingPosition Coord2D) int {
	directionVectors := []Coord2D{
		{-1, 0},  // N
		{-1, 1},  // NE
		{0, 1},   // E
		{1, 1},   // SE
		{1, 0},   // S
		{1, -1},  // SW
		{0, -1},  // W
		{-1, -1}, // NW
	}

	target := "XMAS"
	count := 0

	for _, directionVector := range directionVectors {
		isTarget := true

		for i := 1; i < 4; i++ {
			currentPosition := Coord2D{
				startingPosition.x + directionVector.x*i,
				startingPosition.y + directionVector.y*i,
			}

			if isOutOfBounds(currentPosition, len(matrix), len(matrix[0])-1) {
				isTarget = false
				break
			}

			if matrix[currentPosition.x][currentPosition.y] != target[i] {
				isTarget = false
				break
			}
		}

		if isTarget {
			count++
		}
	}

	return count
}

func (solver *DayFourSolverP1) Solve() []byte {
	var indexes []Coord2D = findIndexes(solver.matrix, 'X')

	result := 0
	for _, index := range indexes {
		result += countXmas(solver.matrix, index)
	}

	return []byte(strconv.Itoa(result))
}

type DayFourSolverP2 struct {
	BaseDayFourSolver
}

func isCrossMas(matrix []string, startingPosition Coord2D) bool {
	mainDiagDirs := []Coord2D{
		{-1, -1}, // NW
		{1, 1},   // SE
	}
	secondDiagDirs := []Coord2D{
		{-1, 1}, // NE
		{1, -1}, // SW
	}
	getDiagChars := func(directionVectors []Coord2D) string {
		chars := make([]rune, 0, 2)

		for _, directionVector := range directionVectors {
			currentPosition := Coord2D{
				startingPosition.x + directionVector.x,
				startingPosition.y + directionVector.y,
			}

			if !isOutOfBounds(currentPosition, len(matrix), len(matrix[0])-1) {
				chars = append(chars, rune(matrix[currentPosition.x][currentPosition.y]))
			}
		}

		return string(chars)
	}

	mainDiag := getDiagChars(mainDiagDirs)
	secondDiag := getDiagChars(secondDiagDirs)

	isValid := func(diag string) bool {
		return diag == "MS" || diag == "SM"
	}

	return isValid(mainDiag) && isValid(secondDiag)
}

func (solver *DayFourSolverP2) Solve() []byte {
	var indexes []Coord2D = findIndexes(solver.matrix, 'A')

	result := 0
	for _, index := range indexes {
		if isCrossMas(solver.matrix, index) {
			result++
		}
	}

	return []byte(strconv.Itoa(result))
}
