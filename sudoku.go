package sudoku

import (
	"math"
)

type Sudoku struct {
	board [][]int
	size  int
}

func NewSudoku(size int) *Sudoku {
	s := &Sudoku{}
	s.size = size
	s.board = make([][]int, size)
	for index, _ := range s.board {
		s.board[index] = make([]int, size)
	}
	return s

}

func validateSequence(seq []int) bool {
	numMap := map[int]bool{}
	for i := 1; i <= len(seq); i++ {
		numMap[i] = false
	}
	for _, value := range seq {
		if value == 0 {
			continue
		} else if numMap[value] {
			return false
		} else {
			numMap[value] = true
		}
	}
	return true

}

func (s Sudoku) validateCol(colInx int) bool {
	col := make([]int, s.size)
	for index, value := range s.board {
		col[index] = value[colInx]
	}
	return validateSequence(col)

}

func (s Sudoku) validateRow(row int) bool {
	return validateSequence(s.board[row])
}

func (s Sudoku) validateBlock(row int, col int) bool {
	block := make([]int, s.size)
	index := 0
	root := int(math.Sqrt(float64(s.size)))
	for i := (row / root) * root; i < ((row/root)+1)*root; i++ {
		for j := (col / root) * root; i < ((col/root)+1)*root; i++ {
			block[index] = s.board[i][j]
			index++
		}
	}
	return validateSequence(block)
}

func (s Sudoku) Solve() {

}

func (s Sudoku) Print() {
	for _, row := range s.board {
		for _, value := range row {
			print(value)
			print(" ")
		}
		println()
	}
}
