package solver

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type BaseDayOneSolver struct {
	leftList  []int
	rightList []int
}

func (solver *BaseDayOneSolver) ReadData(data []byte) error {
	stringData := string(data)
	formattedStringsData := strings.Fields(stringData)

	if len(formattedStringsData)%2 != 0 {
		return fmt.Errorf("invalid input: len(data)&1")
	}

	for i := 1; i < len(formattedStringsData); i += 2 {
		rightListItem, err := strconv.Atoi(formattedStringsData[i])
		if err != nil {
			return fmt.Errorf("failed to parse rightListItem[%d]: %v", i, err)
		}

		leftListItem, err := strconv.Atoi(formattedStringsData[i-1])
		if err != nil {
			return fmt.Errorf("failed to parse leftListItem[%d]: %v", i+1, err)
		}

		solver.rightList = append(solver.rightList, rightListItem)
		solver.leftList = append(solver.leftList, leftListItem)
	}

	return nil
}

type DayOneSolverP1 struct {
	BaseDayOneSolver
}

func (solver *DayOneSolverP1) Solve() []byte {
	sort.Sort(sort.Reverse(sort.IntSlice(solver.leftList)))
	sort.Sort(sort.Reverse(sort.IntSlice(solver.rightList)))

	result := 0
	for i := 0; i < len(solver.leftList); i++ {
		distance := IAbs(solver.leftList[i] - solver.rightList[i])

		result += distance
	}

	return []byte(strconv.Itoa(result))
}

type DayOneSolverP2 struct {
	BaseDayOneSolver
}

func (solver *DayOneSolverP2) Solve() []byte {
	frequencyMap := make(map[int]int)
	for i := 0; i < len(solver.rightList); i++ {
		frequencyMap[solver.rightList[i]]++
	}

	result := 0
	for i := 0; i < len(solver.leftList); i++ {
		currentValue := solver.leftList[i]
		similarityScore := currentValue * frequencyMap[currentValue]

		result += similarityScore
	}

	return []byte(strconv.Itoa(result))
}
