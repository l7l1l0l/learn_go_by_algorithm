package sudoku

import (
	"fmt"
	"github.com/henrylee2cn/faygo"
	"math/rand"
	"time"
)

type SudokuData struct {
	Data     [9][9]int
	MutexNum [9 * 9][]int //互斥的number序列
	TryNum   [9 * 9][]int //尝试过的number序列
}

func (s *SudokuData) Init() {
	for i := 0; i < 9; i++ {
		s.Data[i] = [9]int{}
	}
}

func (s *SudokuData) Print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", s.Data[i][j])
		}
		fmt.Println()
	}
}

func (s *SudokuData) countMutexSlice(i, j int) {
	if i == 0 {
		panic("i == 0") //第一行是随机生成的
	}

	left, top := i-i%3, j-j%3
	index := i*9 + j
	//arr := s.MutexNum[]
	for x := 0; x < left; x++ {
		for y := 0; y < top; y++ {
			s.MutexNum[index] = append(s.MutexNum[index], s.Data[x][y])
		}
	}

	xrange := 3 - i%3 + i
	for x := left; x < j; x++ {
		for y := top; y < xrange; y++ {
			s.MutexNum[index] = append(s.MutexNum[index], s.Data[x][y])
		}
	}
	fmt.Println(s.MutexNum[index])
	time.Sleep(time.Second)
}

func checkNumber(s []int, number int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == number {
			return false
		}
	}
	return true
}

//取一个数字
func (s *SudokuData) getNumber(i, j int) int {
	count := 0
	for {
		if count > 100 {
			faygo.Panicf("rand loop count > 100 [%d,%d]", i, j)
		}
		number := rand.Intn(9) + 1
		count++
		//不能取互斥的数据
		if !checkNumber(s.MutexNum[i*9+i], number) {
			continue
		}
		//不能取已经尝试过的数字
		if !checkNumber(s.TryNum[i*9+i], number) {
			continue
		}
		return number
	}
	return 0
}

//随机生成第一行
func (s *SudokuData) createFirstLine() {
	first := rand.Perm(9)
	for i := 0; i < 9; i++ {
		s.Data[0][i] += first[i] + 1
	}
}

//创建数独
func (s *SudokuData) Create() {
	s.createFirstLine()
	for i := 1; i < 9; i++ {
		for j := 0; j < 9; {
			index := i*9 + i
			s.countMutexSlice(i, j) //计算互斥数据
			if len(s.TryNum[index])+len(s.MutexNum[index]) >= 9 {
				//回退到上一格
				s.clearTmpNums(i, j)
				if j == 0 { //退回上一行
					i-- //退到一行
					j = 8
					s.moveData2Try(i, j)
					break
				} else {
					s.moveData2Try(i, j)
					continue
				}

			} else {
				num := s.getNumber(i, j) //取新数据
				fmt.Println(i, j, s.MutexNum[index], s.TryNum[index], num)
				s.Data[i][j] = num //赋值
				s.addTryNum(i, j, num)
				j++
			}
		}
	}
}

//添加尝试过的数字
func (s *SudokuData) addTryNum(i, j, num int) {
	index := i*9 + j
	s.TryNum[index] = append(s.TryNum[index], num)
}

//将数据区的数字移动到尝试区
func (s *SudokuData) moveData2Try(i, j int) {
	index := i*9 + i
	s.TryNum[index] = append(s.TryNum[index], s.Data[i][j])
	s.Data[i][j] = 0
}

func (s *SudokuData) clearTmpNums(i, j int) {
	index := i*9 + i
	s.MutexNum[index] = nil
	s.TryNum[index] = nil
}

//检查该位置是否合法
func (s *SudokuData) check(i, j, number int) bool {

	return true
}
