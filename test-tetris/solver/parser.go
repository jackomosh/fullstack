package solver

import (
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

type Tetromino struct {
	ID     rune
	Blocks []Point
	H, W   int
}

// ParseFile reads and validates the input file
func ParseFile(filepath string) ([]Tetromino, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Standardize line endings
	text := strings.ReplaceAll(string(content), "\r\n", "\n")
	if len(text) == 0 {
		return nil, os.ErrInvalid
	}

	chunks := strings.Split(text, "\n\n")
	if len(chunks) > 26 { // Maximum limit for English alphabet IDs
		return nil, os.ErrInvalid
	}

	var tetrominoes []Tetromino
	currentID := 'A'

	for _, chunk := range chunks {
		chunk = strings.TrimSpace(chunk)
		if chunk == "" {
			continue
		}

		lines := strings.Split(chunk, "\n")
		if len(lines) != 4 {
			return nil, os.ErrInvalid
		}

		var blocks []Point
		for y, line := range lines {
			if len(line) != 4 {
				return nil, os.ErrInvalid
			}
			for x, char := range line {
				if char == '#' {
					blocks = append(blocks, Point{X: x, Y: y})
				} else if char != '.' {
					return nil, os.ErrInvalid
				}
			}
		}

		if len(blocks) != 4 || !isValidTetromino(blocks) {
			return nil, os.ErrInvalid
		}

		tetrominoes = append(tetrominoes, shiftToOrigin(blocks, currentID))
		currentID++
	}

	if len(tetrominoes) == 0 {
		return nil, os.ErrInvalid
	}

	return tetrominoes, nil
}

// Verifies if the 4 '#' blocks are physically interconnected
func isValidTetromino(blocks []Point) bool {
	connections := 0
	for i := 0; i < len(blocks); i++ {
		for j := i + 1; j < len(blocks); j++ {
			dx := abs(blocks[i].X - blocks[j].X)
			dy := abs(blocks[i].Y - blocks[j].Y)
			if (dx == 1 && dy == 0) || (dx == 0 && dy == 1) {
				connections++
			}
		}
	}
	return connections == 3 || connections == 4
}

// Shifts block coordinates to (0,0) relative origin to allow grid testing alignment
func shiftToOrigin(blocks []Point, id rune) Tetromino {
	minX, minY := 4, 4
	maxX, maxY := -1, -1
	for _, b := range blocks {
		if b.X < minX { minX = b.X }
		if b.Y < minY { minY = b.Y }
	}

	shifted := make([]Point, 4)
	for i, b := range blocks {
		shifted[i] = Point{X: b.X - minX, Y: b.Y - minY}
		if shifted[i].X > maxX { maxX = shifted[i].X }
		if shifted[i].Y > maxY { maxY = shifted[i].Y }
	}

	return Tetromino{
		ID:     id,
		Blocks: shifted,
		W:      maxX + 1,
		H:      maxY + 1,
	}
}

func abs(v int) int {
	if v < 0 { return -v }
	return v
}