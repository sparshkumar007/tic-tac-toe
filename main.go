package main

import (
	"tic-tac-toe/controllers"
	"tic-tac-toe/db"
	"tic-tac-toe/repositories"
	service "tic-tac-toe/service/games"
)

func main() {
	// connect to database (this declares a global variable in db package which has the db connection in it)
	db.ConnectDatabase()

	gameRepo := repositories.NewGameRepository(db.DB)

	gameService := service.NewGameService(gameRepo)

	_ = controllers.NewGameController(gameService)

}
