package main

import (
	"../../sudoku"
)

func main() {
	s := sudoku.SudokuData{}
	s.Create()
	s.Print()
}
