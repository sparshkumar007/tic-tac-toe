package service

import "tic-tac-toe/repositories"

func checkWinner(moves []repositories.GameMoves) (string, bool) {
	board := [3][3]string{}
	for _, move := range moves {
		board[move.PositionX][move.PositionY] = move.PlayerId
	}

	// Check rows and columns
	for i := 0; i < 3; i++ {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0], true
		}
		if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i], true
		}
	}

	// Check diagonals
	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0], true
	}
	if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2], true
	}

	// Check for draw or ongoing game
	if len(moves) == 9 {
		return "", true
	}
	return "", false
}
