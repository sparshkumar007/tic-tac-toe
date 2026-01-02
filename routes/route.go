package routes

import (
	"fmt"
	"tic-tac-toe/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, gameController controllers.GameController) {
	game := r.Group("/game")

	fmt.Print("request has entered here- router \n")
	{
		game.GET("/:id", gameController.GetGame)
		game.POST("/", gameController.CreateGame)
		game.POST("/move", gameController.AddGameMove)
		game.GET("/player/games", gameController.GetAllGamesForPlayer)
		game.GET("/board/:id", gameController.GetGameBoard)
	}
}
