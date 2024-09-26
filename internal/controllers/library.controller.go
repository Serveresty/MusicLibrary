package controllers

import "github.com/jackc/pgx/v4"

type LibraryController struct {
	DB *pgx.Conn
}

func NewLibraryController(DB *pgx.Conn) LibraryController {
	return LibraryController{DB}
}
