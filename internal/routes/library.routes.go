package routes

import (
	"MusicLibrary/internal/controllers"

	"github.com/gin-gonic/gin"
)

type LibraryRouteController struct {
	libraryController controllers.LibraryController
}

func NewLibraryRouteController(libraryController controllers.LibraryController) LibraryRouteController {
	return LibraryRouteController{libraryController}
}

func (lc *LibraryRouteController) LibraryRoute(router *gin.Engine) {

}
