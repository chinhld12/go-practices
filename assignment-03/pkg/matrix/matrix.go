package matrix

import (
	"math"
)

// TraverseRectangles is a recursive function that traverses the matrix and returns a list of visited pairs
func TraverseRectangles(board [][]int, i, j int, visited [][]bool, visitedPairs *[][]int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == 0 || visited[i][j] {
		return
	}

	visited[i][j] = true
	*visitedPairs = append(*visitedPairs, []int{i, j}) // Append (i, j) to the visitedPairs list

	// Recursively check the adjacent cells
	TraverseRectangles(board, i-1, j, visited, visitedPairs) // Up
	TraverseRectangles(board, i+1, j, visited, visitedPairs) // Down
	TraverseRectangles(board, i, j+1, visited, visitedPairs) // Right
	TraverseRectangles(board, i, j-1, visited, visitedPairs) // Left
}

// FindMinRectPoints calculates the minimum and maximum values of i and j in the visitedPairs list
func FindMinRectPoints(visitedPairs [][]int) (lowestI, highestI, lowestJ, highestJ int) {
	// Initialize the variables to the maximum and minimum possible values
	lowestI, highestI = math.MaxInt32, math.MinInt32
	lowestJ, highestJ = math.MaxInt32, math.MinInt32

	// Iterate through the visitedPairs list and update the min/max values
	for _, pair := range visitedPairs {
		i, j := pair[0], pair[1]
		if i < lowestI {
			lowestI = i
		}
		if i > highestI {
			highestI = i
		}
		if j < lowestJ {
			lowestJ = j
		}
		if j > highestJ {
			highestJ = j
		}
	}

	return lowestI, highestI, lowestJ, highestJ
}

// CreateBoolMatrix creates a 2D integer matrix with the given number of rows and columns
func CreateBoolMatrix(rows, cols int) [][]bool {
	initialMatrix := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		initialMatrix[i] = make([]bool, cols)
	}
	return initialMatrix
}
