package service

import "MusicLibrary/internal/repository"

type LibraryService struct {
	repo *repository.LibraryRepository
}

func NewLibraryService(repo *repository.LibraryRepository) *LibraryService {
	return &LibraryService{repo: repo}
}

func (ls *LibraryService) GetMoreInfo() {

}
