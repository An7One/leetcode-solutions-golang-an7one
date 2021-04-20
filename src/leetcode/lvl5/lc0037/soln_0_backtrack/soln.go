// this it NOT yet correct solution
// mostly because of type conversions or comparision between rune and byte
package lc0037

const EMPTY byte = '.'
const RANGE = 9
const DIGITS = "123456789"

func solveSudoku(board [][]byte) {
	backtrack(0, 0, board)
}

func backtrack(curRow int, curCol int, board [][]byte) bool {
	if curCol == RANGE {
		return backtrack(1+curRow, 0, board)
	}

	if curRow == RANGE {
		return true
	}

	for r := curRow; r < RANGE; r++ {
		for c := curCol; c < RANGE; c++ {
			if board[r][c] != EMPTY {
				return backtrack(r, c+1, board)
			}

			for i := range DIGITS {
				digit := DIGITS[i]
				if !isValid(r, c, digit, board) {
					continue
				}

				board[r][c] = digit
				if backtrack(r, c+1, board) {
					return true
				}
				board[r][c] = EMPTY
			}
		}
	}

	return false
}

func isValid(r int, c int, ch byte, board [][]byte) bool {
	for i := 0; i < RANGE; i++ {
		if board[r][i] == ch {
			return false
		}

		if board[i][c] == ch {
			return false
		}

		if board[r/3*3+i/3][c/3*3+i%3] == ch {
			return false
		}
	}

	return true
}
