package sudoku

import (
	"fmt"
	"testing"
)

/*func TestSudokuData_Init(t *testing.T) {
	fmt.Println(rand.Perm(9))
	s := SudokuData{}
	s.Init()
	s.Print()
}*/

func TestSudokuData_Create(t *testing.T) {
	for i := 0; i < 9; i++ {
		fmt.Println(3 - i%3 + i)
	}
	return
	s := SudokuData{}
	s.Init()
	s.Create()
	s.Print()

}
