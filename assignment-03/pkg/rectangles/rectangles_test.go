package rectangles

import (
	"fmt"
	"testing"
)

func catchError() {
		if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
		}
}

func TestCountRectangles(t *testing.T) {
	board := [][]int{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1},
		{1, 1, 0, 1, 1},
	}

	expectedCount := 3

	count := CountRects(board)

	if count != expectedCount {
		t.Errorf("CountRects did not produce the expected count.\nExpected: %d\nGot: %d", expectedCount, count)
	}
}

func TestCountRectanglesEmptyBoard(t *testing.T) {
	var board [][]int

	expectedCount := 0

	defer catchError()

	count := CountRects(board)

	if count != expectedCount {
		t.Errorf("CountRects did not produce the expected count for an empty board.\nExpected: %d\nGot: %d", expectedCount, count)
	}
}

func TestCountRectanglesNoRectangles(t *testing.T) {
	board := [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{1, 1, 1},
	}

	expectedCount := 0

	count := CountRects(board)

	if count != expectedCount {
		t.Errorf("CountRects did not produce the expected count for a board with no rectangles.\nExpected: %d\nGot: %d", expectedCount, count)
	}
}