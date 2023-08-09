package main

import (
	"fmt"

	s "github.com/chinhld12/assignment-03/pkg/rectangles"
)

func main() {
	board := [][]int{
		{1, 0, 1, 1, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	squares := s.CountRects(board)
	fmt.Println("Number of squares on the board:", squares)
}
