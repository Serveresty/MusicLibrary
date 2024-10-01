package controllers

import (
	"MusicLibrary/internal/service"
	"MusicLibrary/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LibraryController struct {
	libService *service.LibraryService
}

func NewLibraryController(libService *service.LibraryService) *LibraryController {
	return &LibraryController{libService: libService}
}

// @Summary Добавить новую песню
// @Description Добавляет новую песню в библиотеку с дополнительной информацией, полученной из внешнего API
// @Tags Песни
// @Accept  json
// @Produce  json
// @Param song body models.SongRequest true "Данные песни"
// @Success 200 {string} string "Ok"
// @Failure 400 {string} string "Неверные данные запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /songs [post]
func (lc *LibraryController) Create(c *gin.Context) {
	var sr models.SongRequest
	if err := c.ShouldBindJSON(&sr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": err.Error()})
		return
	}

	sd, status, err := lc.libService.GetMoreInfo(sr)
	if status != http.StatusOK {
		c.JSON(status, gin.H{"description": err.Error()})
		return
	}

	err = lc.libService.Repo.Create(sr, sd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": "Ok"})
}

// @Summary Получить библиотеку песен
// @Description Возвращает список песен с поддержкой пагинации
// @Tags Песни
// @Accept  json
// @Produce  json
// @Param starts query string false "Номер первого элемента для пагинации"
// @Param limit query string false "Количество элементов для возврата"
// @Success 200 {string} string "Успешный ответ с сообщением 'no data found'"
// @Success 200 {array} models.Song "Успешный ответ с песнями"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /songs [get]
func (lc *LibraryController) GetSongsLibrary(c *gin.Context) {
	starts := c.Query("starts")
	limit := c.Query("limit")

	songs, err := lc.libService.Repo.GetSongsLibrary(starts, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": err.Error()})
		return
	}

	if len(songs) == 0 {
		c.JSON(http.StatusOK, gin.H{"description": "Ok", "content": "no data found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": "Ok", "content": songs})
}

// @Summary Получить текст песни
// @Description Возвращает текст песни по ее ID с поддержкой пагинации (начало и количество строк текста)
// @Tags Песни
// @Accept  json
// @Produce  json
// @Param id path string true "ID песни"
// @Param starts query string true "Номер первой строки для пагинации"
// @Param limit query string true "Количество строк для возврата"
// @Success 200 {array} string "Успешный ответ с текстом песни"
// @Failure 400 {string} string "Неверный формат параметров запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /songs/{id}/text [get]
func (lc *LibraryController) GetSongText(c *gin.Context) {
	songID := c.Param("id")
	starts := c.Query("starts")
	limit := c.Query("limit")

	startINT, err := strconv.Atoi(starts)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "wrong starts num"})
		return
	}
	limitINT, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "wrong limit num"})
		return
	}

	texts, err := lc.libService.Repo.GetSongText(songID, startINT, limitINT)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": "Ok", "content": texts})
}

// @Summary Обновить данные песни
// @Description Обновляет информацию о песне по ее ID
// @Tags Песни
// @Accept  json
// @Produce  json
// @Param id path string true "ID песни"
// @Param song body models.Song true "Обновленные данные песни"
// @Success 200 {string} string "Успешное обновление"
// @Failure 400 {string} string "Неверные данные запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /songs/{id} [patch]
func (lc *LibraryController) Update(c *gin.Context) {
	songID := c.Param("id")
	var updSong models.Song
	if err := c.ShouldBindJSON(&updSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": err.Error()})
		return
	}

	err := lc.libService.Repo.Update(songID, updSong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": "Ok"})
}

// @Summary Удалить песню
// @Description Удаляет песню из библиотеки по ее ID
// @Tags Песни
// @Produce  json
// @Param id path string true "ID песни"
// @Success 200 {string} string "Песня успешно удалена"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /songs/{id} [delete]
func (lc *LibraryController) Delete(c *gin.Context) {
	songID := c.Param("id")
	err := lc.libService.Repo.Delete(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": "Ok"})
}
