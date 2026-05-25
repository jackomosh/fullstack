package main

import (
	"fmt"
	"os"
	"tetris-optimizer/solver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [path_to_file]")
		return
	}

	tetrominoes, err := solver.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	resultBoard := solver.Solve(tetrominoes)
	fmt.Print(resultBoard.String())
}
