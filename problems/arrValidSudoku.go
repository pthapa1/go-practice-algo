package problems

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

- Each row must contain the digits 1-9 without repetition.
- Each column must contain the digits 1-9 without repetition.
- Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
*/

func ValidSudoku(sudokuBoard [][]string) bool {
	if len(sudokuBoard) != 9 {
		return false
	}
	// first verify that the rows and columns are unique
	for i := 0; i < 9; i++ {
		validRow := make(map[string]bool)
		validCol := make(map[string]bool)
		for j := 0; j < 9; j++ {
			rowItem := sudokuBoard[i][j]
			colItem := sudokuBoard[j][i]

			// Check row
			if rowItem != "." {
				if _, exists := validRow[rowItem]; exists {
					return false
				}
				validRow[rowItem] = true
			}

			// Check column
			if colItem != "." {
				if _, exists := validCol[colItem]; exists {
					return false
				}
				validCol[colItem] = true
			}
		}
	}

	// Now, check whether 3x3 is unique
	// first divide the wider 9 by 9 as if we have an array of 3 x 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			validGrid := make(map[string]bool)
			// now loop insdie the each grid's value and add each item to the validGrid
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					eachItemOfSubGrid := sudokuBoard[i*3+k][j*3+l]

					if eachItemOfSubGrid != "." {
						if _, exists := validGrid[eachItemOfSubGrid]; exists {
							return false
						}
						validGrid[eachItemOfSubGrid] = true
					}
				}
			}
		}
	}

	return true
}
