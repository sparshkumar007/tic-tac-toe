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
	CreateGame(c *gin.Context)
	AddGameMove(c *gin.Context)
	GetAllGamesForPlayer(c *gin.Context)
	GetGameBoard(c *gin.Context)
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

func (ctl *gameController) CreateGame(c *gin.Context) {
	fmt.Print("request has reached here - userController")
	req, err := mappers.DecodeCreateGameRequest(c)
	if err != nil {
		c.JSON(400, helpers.JsonResp{
			Message: err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	resp, err := ctl.GameService.CreateGame(ctx, req)
	if err != nil {
		c.JSON(400, helpers.JsonResp{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, helpers.JsonResp{
		Data:    resp,
		Message: "game created successfully",
	})
}

func (ctl *gameController) AddGameMove(c *gin.Context) {
	fmt.Print("request has reached here - userController")
	req, err := mappers.DecodeAddGameMoveRequest(c)
	if err != nil {
		c.JSON(400, helpers.JsonResp{
			Message: err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	resp, err := ctl.GameService.AddGameMove(ctx, req)
	if err != nil {
		c.JSON(400, helpers.JsonResp{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, helpers.JsonResp{
		Data:    resp,
		Message: "game move added successfully",
	})
}

func (ctl *gameController) GetAllGamesForPlayer(c *gin.Context) {
	fmt.Print("request has reached here - userController")
	emailId := c.Query("email_id")
	if emailId == "" {
		c.JSON(400, helpers.JsonResp{
			Message: "email_id query parameter is required",
		})
		return
	}

	ctx := c.Request.Context()
	resp, err := ctl.GameService.GetAllGamesForPlayer(ctx, emailId)
	if err != nil {
		c.JSON(400, helpers.JsonResp{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, helpers.JsonResp{
		Data:    resp,
		Message: "games fetched successfully",
	})
}

func (ctl *gameController) GetGameBoard(c *gin.Context) {
	fmt.Print("request has reached here - userController")
	req, err := mappers.DecodeGetGameBoardRequest(c)
	if err != nil {
		c.JSON(400, helpers.JsonResp{
			Message: err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	resp, err := ctl.GameService.GetGameBoard(ctx, req)
	if err != nil {
		c.JSON(400, helpers.JsonResp{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, helpers.JsonResp{
		Data:    resp,
		Message: "game board fetched successfully",
	})
}
