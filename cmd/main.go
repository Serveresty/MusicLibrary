package main

import (
	"MusicLibrary/configs"
	"MusicLibrary/database"
	"MusicLibrary/internal/controllers"
	"MusicLibrary/internal/routes"
	"MusicLibrary/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dbConf := configs.LoadDBConfig()
	serverConf := configs.LoadServerConfig()

	dbConn, err := database.DBInit(dbConf)
	if err != nil {
		log.Fatalf("error: db connection err, %v", err)
	}

	loggers := logger.NewLoggers()
	libController := controllers.NewLibraryController(dbConn, loggers)
	libRController := routes.NewLibraryRouteController(libController)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	libRController.LibraryRoute(router)

	libController.Logs.Info.Printf("Server starts on %s", serverConf.Host+":"+serverConf.Port)
	if err := router.Run(serverConf.Host + ":" + serverConf.Port); err != nil {
		log.Fatalf("error: server didn't run, %v", err)
	}
}
