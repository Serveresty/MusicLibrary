package service

import (
	"MusicLibrary/internal/repository"
	"MusicLibrary/models"
	"encoding/json"
	"fmt"
	"log/slog"
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
	debugLogg := func(url string, status int, err string, content any) {
		ls.Repo.Logs.DebugLog(
			"Get detail song's info request",
			slog.String("method", "GET"),
			slog.String("path", url),
			slog.Int("status", status),
			slog.String("error", err),
			slog.Any("content", content),
		)
	}
	api := os.Getenv("API_URL")
	url := fmt.Sprintf(api+"?group=%s&song=%s", url.QueryEscape(info.Group), url.QueryEscape(info.Song))

	resp, err := http.Get(url)
	if err != nil {
		debugLogg(url, resp.StatusCode, err.Error(), nil)
		return models.SongDetail{}, resp.StatusCode, err
	}
	defer resp.Body.Close()

	var songDetail models.SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		debugLogg(url, http.StatusBadRequest, err.Error(), nil)
		return models.SongDetail{}, http.StatusBadRequest, err
	}

	debugLogg(url, http.StatusOK, "", songDetail)
	return songDetail, http.StatusOK, nil
}
