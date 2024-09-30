package repository

import (
	"MusicLibrary/models"
	"MusicLibrary/pkg/logger"
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
)

type LibraryRepository struct {
	db   *pgx.Conn
	Logs *logger.Loggers
}

func NewLibraryRepository(db *pgx.Conn, loggers *logger.Loggers) *LibraryRepository {
	return &LibraryRepository{db: db, Logs: loggers}
}

func (lc *LibraryRepository) Create(sr models.SongRequest, sd models.SongDetail) error {
	_, err := lc.db.Exec(context.Background(), `INSERT INTO "songs" ("group", "song", "release_date", "text", "link") VALUES ($1,$2,$3,$4,$5)`,
		sr.Group, sr.Song, sd.ReleaseDate, sd.Text, sd.Link)
	if err != nil {
		return err
	}
	return nil
}

func (lc *LibraryRepository) GetSongsLibrary(minID, limit string) ([]models.Song, error) {
	rows, err := lc.db.Query(context.Background(), `SELECT * FROM "songs" WHERE id >= $1 LIMIT $2`, minID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	return songs, nil
}

func (lc *LibraryRepository) GetSongText(songID string, starts, limit int) ([]string, error) {
	rows, err := lc.db.Query(context.Background(), `SELECT "text" FROM "songs" WHERE id = $1 `, songID)
	if err != nil {
		return nil, err
	}

	var text string
	for rows.Next() {
		if err := rows.Scan(&text); err != nil {
			return nil, err
		}
	}

	texts := strings.Split(text, "\n\n")
	if starts > len(texts)-1 {
		return nil, fmt.Errorf("invalid input value, max starts value is %v", len(texts)-1)
	}
	if starts < 0 || limit < 0 {
		return nil, fmt.Errorf("invalid input value, below zero")
	}
	if len(texts)-1 < (starts + limit) {
		return texts[starts:], nil
	}

	return texts[starts : starts+limit], nil
}

func (lc *LibraryRepository) Update(songID string, updSong models.Song) error {
	query := `
        UPDATE songs
        SET 
			"group" = COALESCE(NULLIF($1, ''), "group"),
			"song" = COALESCE(NULLIF($2, ''), "song"),
			"release_date" = COALESCE(NULLIF($3, ''), "release_date"),
            "text" = COALESCE(NULLIF($4, ''), "text"),
            "link" = COALESCE(NULLIF($5, ''), "link")
        WHERE id = $6;
    `

	_, err := lc.db.Exec(context.Background(), query, updSong.Group, updSong.Song, updSong.ReleaseDate, updSong.Text, updSong.Link, songID)
	if err != nil {
		return err
	}

	return nil
}

func (lc *LibraryRepository) Delete(songID string) error {
	_, err := lc.db.Exec(context.Background(), `DELETE FROM "songs" WHERE id = $1`, songID)
	if err != nil {
		return err
	}

	return nil
}
