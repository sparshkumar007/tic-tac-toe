package controllers

import (
	"fmt"
	"tic-tac-toe/helpers"
	"tic-tac-toe/mappers"
	service "tic-tac-toe/service/games"

	"github.com/gin-gonic/gin"
)

type GameController interface {
	GetGame(c *gin.Context)
}

type gameController struct {
	GameService service.GameService
}

func NewGameController(gameService service.GameService) GameController {
	return &gameController{
		GameService: gameService,
	}
}

func (ctl *gameController) GetGame(c *gin.Context) {
	fmt.Print("request has reached here - userController")
	req, err := mappers.DecodeGetGameRequest(c)
	if err != nil {
		c.JSON(400, helpers.JsonResp{
			Message: err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	resp, err := ctl.GameService.GetGame(ctx, req)
	if err != nil {
		c.JSON(400, helpers.JsonResp{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, helpers.JsonResp{
		Data:    resp,
		Message: "game fetched successfully",
	})
}
