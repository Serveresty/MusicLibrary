package routes

import (
	"MusicLibrary/internal/controllers"

	_ "MusicLibrary/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type LibraryRouteController struct {
	libContr *controllers.LibraryController
}

func NewLibraryRouteController(libraryController *controllers.LibraryController) *LibraryRouteController {
	return &LibraryRouteController{libraryController}
}

func (lc *LibraryRouteController) LibraryRoute(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	songs := router.Group("songs")
	{
		songs.GET("/", lc.libContr.GetSongsLibrary)
		songs.GET("/:id/text", lc.libContr.GetSongText)

		songs.DELETE("/:id", lc.libContr.Delete)

		songs.PATCH("/:id", lc.libContr.Update)

		songs.POST("/", lc.libContr.Create)
	}
}
