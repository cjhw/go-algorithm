package Recall

func solveSudoku(board [][]byte) {
	var isVaild func(row, col int, k byte, board [][]byte) bool

	isVaild = func(row, col int, k byte, board [][]byte) bool {
		for i := 0; i < 9; i++ {
			if board[row][i] == k {
				return false
			}
		}

		for i := 0; i < 9; i++ { //列
			if board[i][col] == k {
				return false
			}
		}

		startRow := (row / 3) * 3
		startCol := (col / 3) * 3

		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == k {
					return false
				}
			}
		}

		return true
	}

	var backtracking func(board [][]byte) bool

	backtracking = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					if isVaild(i, j, byte(k), board) {
						board[i][j] = byte(k)
						if backtracking(board) == true {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}

		return true
	}

	backtracking(board)
}