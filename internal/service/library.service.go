package service

import (
	"MusicLibrary/internal/repository"
	"MusicLibrary/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type LibraryService struct {
	Repo *repository.LibraryRepository
}

func NewLibraryService(repo *repository.LibraryRepository) *LibraryService {
	return &LibraryService{Repo: repo}
}

func (ls *LibraryService) GetMoreInfo(info models.SongRequest) (models.SongDetail, int, error) {
	api := os.Getenv("API_URL")
	url := fmt.Sprintf(api+"?group=%s&song=%s", url.QueryEscape(info.Group), url.QueryEscape(info.Song))

	resp, err := http.Get(url)
	if err != nil {
		return models.SongDetail{}, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	var songDetail models.SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return models.SongDetail{}, http.StatusBadRequest, err
	}

	return songDetail, http.StatusOK, nil
}
