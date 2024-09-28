package routes

import (
	"MusicLibrary/internal/controllers"

	"github.com/gin-gonic/gin"
)

type LibraryRouteController struct {
	libContr *controllers.LibraryController
}

func NewLibraryRouteController(libraryController *controllers.LibraryController) *LibraryRouteController {
	return &LibraryRouteController{libraryController}
}

func (lc *LibraryRouteController) LibraryRoute(router *gin.Engine) {
	songs := router.Group("songs")
	{
		songs.GET("/", lc.libContr.GetSongsLibrary)
		songs.GET("/:id/text", lc.libContr.GetSongText)

		songs.DELETE("/:id", lc.libContr.Delete)

		songs.PATCH("/:id", lc.libContr.Update)

		songs.POST("/", lc.libContr.Create)
	}
}
