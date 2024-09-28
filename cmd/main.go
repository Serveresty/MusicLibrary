package main

import (
	"MusicLibrary/configs"
	"MusicLibrary/database"
	"MusicLibrary/internal/controllers"
	"MusicLibrary/internal/repository"
	"MusicLibrary/internal/routes"
	"MusicLibrary/internal/service"
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
	libRepo := repository.NewLibraryRepository(dbConn, loggers)
	libServ := service.NewLibraryService(libRepo)
	libCont := controllers.NewLibraryController(libServ)
	libRoute := routes.NewLibraryRouteController(libCont)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	libRoute.LibraryRoute(router)

	libRepo.Logs.Info.Printf("Server starts on %s", serverConf.Host+":"+serverConf.Port)
	if err := router.Run(serverConf.Host + ":" + serverConf.Port); err != nil {
		log.Fatalf("error: server didn't run, %v", err)
	}
}
