package repository

import (
	"MusicLibrary/pkg/logger"

	"github.com/jackc/pgx/v4"
)

type LibraryRepository struct {
	DB   *pgx.Conn
	Logs *logger.Loggers
}

func NewLibraryRepository(db *pgx.Conn, loggers *logger.Loggers) *LibraryRepository {
	return &LibraryRepository{DB: db, Logs: loggers}
}

func (lc *LibraryRepository) Create() {

}

func (lc *LibraryRepository) GetSongsLibrary() {

}

func (lc *LibraryRepository) GetSongText() {

}

func (lc *LibraryRepository) Update() {

}

func (lc *LibraryRepository) Delete() {

}
