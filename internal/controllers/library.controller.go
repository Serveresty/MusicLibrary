package controllers

import (
	"MusicLibrary/pkg/logger"

	"github.com/jackc/pgx/v4"
)

type LibraryController struct {
	DB   *pgx.Conn
	Logs *logger.Loggers
}

func NewLibraryController(db *pgx.Conn, loggers *logger.Loggers) LibraryController {
	return LibraryController{DB: db, Logs: loggers}
}
