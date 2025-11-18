package service

import (
	"tic-tac-toe/helpers"
)

func ValidateGetGameRequest(req GetGameRequest) error {
	if req.GameId == 0 {
		return helpers.BadRequest("invalid game id")
	}
	return nil
}
