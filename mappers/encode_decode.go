package mappers

import (
	service "tic-tac-toe/service/games"

	"github.com/gin-gonic/gin"
)

func DecodeGetGameRequest(c *gin.Context) (service.GetGameRequest, error) {
	var req service.GetGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return service.GetGameRequest{}, err
	}
	return req, nil
}

func DecodeCreateGameRequest(c *gin.Context) (service.CreateGameRequest, error) {
	var req service.CreateGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return service.CreateGameRequest{}, err
	}
	return req, nil
}

func DecodeAddGameMoveRequest(c *gin.Context) (service.AddGameMoveRequest, error) {
	var req service.AddGameMoveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return service.AddGameMoveRequest{}, err
	}
	return req, nil
}

func DecodeGetGameBoardRequest(c *gin.Context) (service.GetGameBoardRequest, error) {
	var req service.GetGameBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return service.GetGameBoardRequest{}, err
	}
	return req, nil
}
