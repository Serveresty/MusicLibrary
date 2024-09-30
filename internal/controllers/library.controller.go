package controllers

import (
	"MusicLibrary/internal/service"
	"MusicLibrary/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LibraryController struct {
	libService *service.LibraryService
}

func NewLibraryController(libService *service.LibraryService) *LibraryController {
	return &LibraryController{libService: libService}
}

func (lc *LibraryController) Create(c *gin.Context) {
	var sr models.SongRequest
	if err := c.ShouldBindJSON(&sr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sd, status := lc.libService.GetMoreInfo(sr)
	if status != 200 {
		c.JSON(status, gin.H{"error": "err"})
		return
	}

	err := lc.libService.Repo.Create(sr, sd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song created successfully"})
}

func (lc *LibraryController) GetSongsLibrary(c *gin.Context) {
	starts := c.Query("starts")
	limit := c.Query("limit")

	songs, err := lc.libService.Repo.GetSongsLibrary(starts, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(songs) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": "no data found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": songs})
}

func (lc *LibraryController) GetSongText(c *gin.Context) {

}

func (lc *LibraryController) Update(c *gin.Context) {

}

func (lc *LibraryController) Delete(c *gin.Context) {
	songID := c.Param("id")
	err := lc.libService.Repo.Delete(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "song has been deleted"})
}
