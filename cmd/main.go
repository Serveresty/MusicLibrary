package main

import (
	"MusicLibrary/configs"
	"MusicLibrary/database"
	"MusicLibrary/internal/controllers"
	"MusicLibrary/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	dbConf := configs.LoadDBConfig()
	serverConf := configs.LoadServerConfig()

	dbConn, err := database.DBInit(dbConf)
	if err != nil {
		log.Fatalf("error: db connection err, %v", err)
	}

	libraryController := controllers.NewLibraryController(dbConn)
	libraryRouteController := routes.NewLibraryRouteController(libraryController)

	router := gin.Default()

	libraryRouteController.LibraryRoute(router)
	if err := router.Run(":" + serverConf.Port); err != nil {
		log.Fatalf("error: server didn't run, %v", err)
	}
}
