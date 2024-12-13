package solver

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type BaseDayTwoSolver struct {
	reports [][]int
}

func (solver *BaseDayTwoSolver) ReadData(data []byte) error {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		values := strings.Fields(line)

		intValues := []int{}
		for _, value := range values {
			parsedValue, err := strconv.Atoi(value)

			if err != nil {
				return fmt.Errorf("failed to parse rightListItem[%q]: %v", value, err)
			}

			intValues = append(intValues, parsedValue)
		}

		solver.reports = append(solver.reports, intValues)
	}

	return nil
}

type DayTwoSolverP1 struct {
	BaseDayTwoSolver
}

func CheckReport(data []int) bool {
	if !sort.IsSorted(sort.IntSlice(data)) && !sort.IsSorted(sort.Reverse(sort.IntSlice(data))) {
		return false
	}

	isSafe := true
	for i := 1; i < len(data); i++ {
		if data[i] == data[i-1] || IAbs(data[i]-data[i-1]) > 3 {
			isSafe = false
			break
		}
	}

	return isSafe
}

func (solver *DayTwoSolverP1) Solve() []byte {
	safeReports := 0

	for _, report := range solver.reports {
		if CheckReport(report) {
			safeReports++
		}
	}

	return []byte(strconv.Itoa(safeReports))
}

type DayTwoSolverP2 struct {
	BaseDayTwoSolver
}

func (solver *DayTwoSolverP2) Solve() []byte {
	safeReports := 0

	for _, report := range solver.reports {
		if CheckReport(report) {
			safeReports++
			continue
		}

		for i := range report {
			firstPart := append([]int(nil), report[:i]...)
			copyReport := append(firstPart, report[i+1:]...)

			if CheckReport(copyReport) {
				safeReports++
				break
			}
		}

	}

	return []byte(strconv.Itoa(safeReports))
}
