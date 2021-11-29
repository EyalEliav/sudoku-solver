package sudoku

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
