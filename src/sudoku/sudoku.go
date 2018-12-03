package sudoku

import (
	"fmt"
	"github.com/henrylee2cn/faygo"
	"math/rand"
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

func (s *SudokuData) addMutexNum(i, j, number int) {
	index := i*9 + j
	for x := 0; x < len(s.MutexNum[index]); x++ {
		if number == s.MutexNum[index][x] {
			return
		}
	}
	s.MutexNum[index] = append(s.MutexNum[index], number)
}

func (s *SudokuData) countMutexSlice(i, j int) {
	if i == 0 {
		panic("i == 0") //第一行是随机生成的
	}

	left, top := i-i%3, j-j%3
	//index := i*9 + j
	//arr := s.MutexNum[]
	/*for x := i; x < top; x++ {
		for y := 0; y <= left; y++ {
			s.MutexNum[index] = append(s.MutexNum[index], s.Data[x][y])
		}
	}*/
	for y := 0; y < top; y++ {
		//s.MutexNum[index] = append(s.MutexNum[index], s.Data[i][y])
		s.addMutexNum(i, j, s.Data[i][y])
	}

	for x := 0; x < left; x++ {
		//s.MutexNum[index] = append(s.MutexNum[index], s.Data[x][j])
		s.addMutexNum(i, j, s.Data[x][j])
	}

	/*xrange := 3 - i%3 + i
	for x := left; x <= i; x++ {
		for y := top; y < xrange; y++ {
			if s.Data[x][y] <= 0 {
				break
			}
			s.MutexNum[index] = append(s.MutexNum[index], s.Data[x][y])
		}
	}*/

	endFlag := false
	for x := left; x < left+3; x++ {
		if endFlag {
			break
		}
		for y := top; y < top+3; y++ {
			if x == i && y == j {
				endFlag = true
				break
			}
			s.addMutexNum(i, j, s.Data[x][y])
			//s.MutexNum[index] = append(s.MutexNum[index], s.Data[x][y])
		}
	}
	//log.Printf("mutex %+v", s.MutexNum[index])
	//time.Sleep(100 * time.Microsecond)
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
	index := i*9 + j
	for {
		if count > 100 {
			faygo.Panicf("rand loop count > 100 [%d,%d]", i, j)
		}
		number := rand.Intn(9) + 1
		count++
		//不能取互斥的数据
		if !checkNumber(s.MutexNum[index], number) {
			continue
		}
		//不能取已经尝试过的数字
		if !checkNumber(s.TryNum[index], number) {
			continue
		}
		//log.Printf("mutex len[%d] try len[%d]", len(s.MutexNum[index]), len(s.TryNum[index]))
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
	//fmt.Printf("first %+v\n", s.Data[0])

}

func (s *SudokuData) PrintCurrent(i int) {
	fmt.Println("------------------------")
	for x := 0; x <= i; x++ {
		fmt.Printf("%+v\n", s.Data[x])
	}
	fmt.Println("------------------------")
}

//创建数独
func (s *SudokuData) Create() {
	s.createFirstLine()
	for i := 1; i < 9; i++ {
		firstError, count := 0, 0

		for j := 0; j < 9; {
			index := i*9 + j
			s.countMutexSlice(i, j) //计算互斥数据
			if len(s.TryNum[index])+len(s.MutexNum[index]) >= 9 {
				if firstError > 0 {
					if index > firstError {
						firstError = index
						count = 1
					}
				} else {
					firstError = index
					count++
				}

				/*fmt.Printf("{%d,%d} mutex %+v try %+v\n", i, j,
				s.MutexNum[index], s.TryNum[index])*/

				//遍历到此处，发现已经无法填值了
				s.clearTmpNums(i, j)

				//s.PrintCurrent(i) //打印当前数据

				//是否要回到上一行
				if j-count < 0 {
					for k := 0; k < j; k++ {
						s.clearTmpNums(i, k)
					}
					i--
					j = 7

				} else {
					for k := 0; k < count; k++ {
						s.clearTmpNums(i, k)
						j--
					}
				}
				continue
			} else {
				num := s.getNumber(i, j) //取新数据
				if num == 0 {
					panic("num is 0")
				}
				s.Data[i][j] = num //赋值
				s.addTryNum(i, j, num)
				//log.Printf("index[%d] {%d,%d} %+v, %+v num[%d]", i*9+j, i, j, s.MutexNum[index], s.TryNum[index], num)
				j++
			}
		}
	}
}

//添加尝试过的数字
func (s *SudokuData) addTryNum(i, j, num int) {
	index := i*9 + j
	//log.Printf("add try num index[%d] num[%d]", index, num)
	s.TryNum[index] = append(s.TryNum[index], num)
}

//将数据区的数字移动到尝试区
func (s *SudokuData) moveData2Try(i, j int) {
	index := i*9 + j
	s.TryNum[index] = append(s.TryNum[index], s.Data[i][j])
	s.Data[i][j] = 0
}

func (s *SudokuData) clearTmpNums(i, j int) {

	index := i*9 + j
	s.MutexNum[index] = nil
	s.TryNum[index] = nil
	//s.Data[i][j] = 0
	/*fmt.Printf("clear {%d,%d} len(mutex[%d])=%d len(tryNum[%d])=%d\n",
	i, j, index, len(s.MutexNum[index]), index, len(s.TryNum[index]))*/
}

//检查该位置是否合法
func (s *SudokuData) SimpleCheck() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for x := 0; x < 9; x++ {
				if x != i && s.Data[i][j] == s.Data[x][j] {
					return false
				}
			}
			for y := 0; y < 0; y++ {
				if y != j && s.Data[i][j] == s.Data[i][y] {
					return false
				}
			}
		}
	}
	return true
}
