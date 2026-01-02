package service

import (
	"tic-tac-toe/helpers"
	"tic-tac-toe/repositories"
)

func ValidateGetGameRequest(req GetGameRequest) error {
	if req.GameId == 0 {
		return helpers.BadRequest("invalid game id")
	}
	return nil
}

func ValidateCreateGameRequest(req CreateGameRequest) error {
	if len(req.EmailIds) != 2 {
		return helpers.BadRequest("exactly two email ids are required")
	}
	return nil
}

func ValidateAddGameMoveRequest(req AddGameMoveRequest) error {
	if req.GameId == 0 {
		return helpers.BadRequest("invalid game id")
	}
	if req.PlayerId == "" {
		return helpers.BadRequest("invalid player id")
	}
	if req.PositionX > 2 {
		return helpers.BadRequest("invalid position x")
	}
	if req.PositionY > 2 {
		return helpers.BadRequest("invalid position y")
	}
	return nil
}

func valiateMove(game repositories.Game, moves []repositories.GameMoves, move repositories.GameMoves) error {
	if game.Status == repositories.GameStatusCompleted {
		return helpers.BadRequest("game is already completed")
	}
	if game.LastMoveBy == move.PlayerId {
		return helpers.BadRequest("not your turn to play")
	}
	for _, m := range moves {
		if m.PositionX == move.PositionX && m.PositionY == move.PositionY {
			return helpers.BadRequest("position already taken")
		}
	}
	return nil
}
