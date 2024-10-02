package main

import (
	"MusicLibrary/configs"
	"MusicLibrary/database"
	"MusicLibrary/internal/controllers"
	"MusicLibrary/internal/repository"
	"MusicLibrary/internal/routes"
	"MusicLibrary/internal/service"
	"MusicLibrary/pkg/logger"
	"context"
	"io"
	"log"
	"log/slog"

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
	defer dbConn.Close(context.Background())

	database.RunMigrations(dbConn)

	loggers, err := logger.NewLoggers()
	if err != nil {
		log.Fatalf("error: init loggers err, %v", err)
	}
	libRepo := repository.NewLibraryRepository(dbConn, loggers)
	libServ := service.NewLibraryService(libRepo)
	libCont := controllers.NewLibraryController(libServ)
	libRoute := routes.NewLibraryRouteController(libCont)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard
	router := gin.Default()

	libRoute.LibraryRoute(router)

	loggers.InfoLog(
		"Server started",
		slog.String("address", serverConf.Host),
		slog.String("port", serverConf.Port),
	)

	if err := router.Run(serverConf.Host + ":" + serverConf.Port); err != nil {
		log.Fatalf("error: server didn't run, %v", err)
	}
}
