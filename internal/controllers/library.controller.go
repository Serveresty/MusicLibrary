package controllers

import (
	"MusicLibrary/internal/service"

	"github.com/gin-gonic/gin"
)

type LibraryController struct {
	libService *service.LibraryService
}

func NewLibraryController(libService *service.LibraryService) *LibraryController {
	return &LibraryController{libService: libService}
}

func (lc *LibraryController) Create(c *gin.Context) {

}

func (lc *LibraryController) GetSongsLibrary(c *gin.Context) {

}

func (lc *LibraryController) GetSongText(c *gin.Context) {

}

func (lc *LibraryController) Update(c *gin.Context) {

}

func (lc *LibraryController) Delete(c *gin.Context) {

}
