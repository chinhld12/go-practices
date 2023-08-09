package rectangles

import "github.com/chinhld12/assignment-03/pkg/matrix"

func isRectangles(board [][]int, lowestI, highestI, lowestJ, highestJ int) bool {
	isRectangle := true
	for i := lowestI; i <= highestI; i++ {
		for j := lowestJ; j <= highestJ; j++ {
			if board[i][j] == 0 {
				isRectangle = false
				break
			}
		}
	}
	return isRectangle
}

// CountRects counts the number of squares in the given matrix
func CountRects(board [][]int) int {
	if len(board) == 0 {
		panic("Empty board")
	}

	if len(board[0]) == 0 {
		panic("Row defined is not correct")
	}

	rows, cols := len(board), len(board[0])
	visited := matrix.CreateBoolMatrix(rows, cols)

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] != 0 && !visited[i][j] {
				var visitedPairs [][]int

				matrix.TraverseRectangles(board, i, j, visited, &visitedPairs)
				lowestI, highestI, lowestJ, highestJ := matrix.FindMinRectPoints(visitedPairs)
				if isRectangles(board, lowestI, highestI, lowestJ, highestJ) {
					count++
				}
			} else if board[i][j] == 0 {
				visited[i][j] = true
			}
		}
	}

	return count
}
