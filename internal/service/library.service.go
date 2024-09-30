package service

import (
	"MusicLibrary/internal/repository"
	"MusicLibrary/models"
)

type LibraryService struct {
	Repo *repository.LibraryRepository
}

func NewLibraryService(repo *repository.LibraryRepository) *LibraryService {
	return &LibraryService{Repo: repo}
}

func (ls *LibraryService) GetMoreInfo(info models.SongRequest) (models.SongDetail, int) {

	return models.SongDetail{}, 200
}
