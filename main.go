package main

import (
	"advent_of_code/solver"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	var solver solver.AocSolver = &solver.DayThreeSolverP2{}

	if err := solver.ReadData(content); err != nil {
		fmt.Println("Error:", err)
	}

	var result []byte = solver.Solve()
	err = os.WriteFile("output.txt", result, 0666)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
