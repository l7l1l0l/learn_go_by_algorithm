package main

import (
	"../../sudoku"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := sudoku.SudokuData{}

	s.Create()

	s.Print()
	secs := time.Since(start).Seconds()

	fmt.Println("check result:", s.SimpleCheck(), " use time:", secs)
}
