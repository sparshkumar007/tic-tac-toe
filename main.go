package main

import (
	"log"
	"os"
	"tic-tac-toe/controllers"
	"tic-tac-toe/db"
	"tic-tac-toe/repositories"
	"tic-tac-toe/routes"
	service "tic-tac-toe/service/games"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to database (this declares a global variable in db package which has the db connection in it)
	db.ConnectDatabase()

	gameRepo := repositories.NewGameRepository(db.DB)

	gameService := service.NewGameService(gameRepo, db.DB)

	gameController := controllers.NewGameController(gameService)

	router := gin.Default()

	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:3001",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	routes.SetupRoutes(router, gameController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on http://localhost:%s", port)
	router.Run(":" + port)

}
