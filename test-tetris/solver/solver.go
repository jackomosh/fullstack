package solver

import (
	"math"
	"strings"
)

type Board [][]rune

func NewBoard(size int) Board {
	b := make(Board, size)
	for i := range b {
		b[i] = make([]rune, size)
		for j := range b[i] {
			b[i][j] = '.'
		}
	}
	return b
}

func Solve(tetrominoes []Tetromino) Board {
	// Start with the minimum possible theoretical square size boundary
	minSize := int(math.Ceil(math.Sqrt(float64(len(tetrominoes) * 4))))
	
	for size := minSize; ; size++ {
		board := NewBoard(size)
		if backtrack(board, tetrominoes, 0) {
			return board
		}
	}
}

func backtrack(board Board, list []Tetromino, index int) bool {
	if index == len(list) {
		return true
	}

	t := list[index]
	size := len(board)

	for y := 0; y <= size-t.H; y++ {
		for x := 0; x <= size-t.W; x++ {
			if canPlace(board, t, x, y) {
				place(board, t, x, y, t.ID)
				if backtrack(board, list, index+1) {
					return true
				}
				place(board, t, x, y, '.') // Backtrack step cleanup
			}
		}
	}
	return false
}

func canPlace(board Board, t Tetromino, x, y int) bool {
	for _, b := range t.Blocks {
		if board[y+b.Y][x+b.X] != '.' {
			return false
		}
	}
	return true
}

func place(board Board, t Tetromino, x, y int, id rune) {
	for _, b := range t.Blocks {
		board[y+b.Y][x+b.X] = id
	}
}

func (b Board) String() string {
	var sb strings.Builder
	for _, row := range b {
		for _, char := range row {
			sb.WriteRune(char)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}