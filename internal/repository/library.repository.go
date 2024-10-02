package repository

import (
	"MusicLibrary/models"
	"MusicLibrary/pkg/logger"
	"context"
	"fmt"
	"log/slog"
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

func (lr *LibraryRepository) Create(sr models.SongRequest, sd models.SongDetail) error {
	debugLogg := func(result, err string) {
		lr.Logs.DebugLog(
			"Create song db method",
			slog.Any("info", sr),
			slog.Any("detail", sd),
			slog.String("error", err),
			slog.String("result", result),
		)
	}

	_, err := lr.db.Exec(context.Background(), `INSERT INTO "songs" ("group", "song", "release_date", "text", "link") VALUES ($1,$2,$3,$4,$5)`,
		sr.Group, sr.Song, sd.ReleaseDate, sd.Text, sd.Link)
	if err != nil {
		debugLogg("failure", err.Error())
		return err
	}

	debugLogg("success", "")
	return nil
}

func (lr *LibraryRepository) GetSongsLibrary(minID, limit string) ([]models.Song, error) {
	debugLogg := func(content any, result, err string) {
		lr.Logs.DebugLog(
			"Get song's library db method",
			slog.Group("params",
				slog.String("minID", minID),
				slog.String("limit", limit),
			),
			slog.String("error", err),
			slog.String("result", result),
			slog.Any("content", content),
		)
	}

	rows, err := lr.db.Query(context.Background(), `SELECT * FROM "songs" WHERE id >= $1 LIMIT $2`, minID, limit)
	if err != nil {
		debugLogg(nil, "failure", err.Error())
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			debugLogg(nil, "failure", err.Error())
			return nil, err
		}

		songs = append(songs, song)
	}

	debugLogg(songs, "success", "")
	return songs, nil
}

func (lr *LibraryRepository) GetSongText(songID string, starts, limit int) ([]string, error) {
	debugLogg := func(content any, result, err string) {
		lr.Logs.DebugLog(
			"Get song's text db method",
			slog.Group("params",
				slog.String("songID", songID),
				slog.Int("starts", starts),
				slog.Int("limit", limit),
			),
			slog.String("error", err),
			slog.String("result", result),
			slog.Any("content", content),
		)
	}

	rows, err := lr.db.Query(context.Background(), `SELECT "text" FROM "songs" WHERE id = $1 `, songID)
	if err != nil {
		debugLogg(nil, "failure", err.Error())
		return nil, err
	}

	var text string
	for rows.Next() {
		if err := rows.Scan(&text); err != nil {
			debugLogg(nil, "failure", err.Error())
			return nil, err
		}
	}

	texts := strings.Split(text, "\n\n")
	if starts > len(texts)-1 {
		err1 := fmt.Errorf("invalid input value, max starts value is %v", len(texts)-1)
		debugLogg(nil, "failure", err1.Error())
		return nil, err1
	}
	if starts < 0 || limit < 0 {
		err1 := fmt.Errorf("invalid input value, below zero")
		debugLogg(nil, "failure", err1.Error())
		return nil, err1
	}
	var result []string
	if len(texts)-1 < (starts + limit) {
		result = texts[starts:]
		debugLogg(result, "success", "")
		return result, nil
	}

	result = texts[starts : starts+limit]
	debugLogg(result, "success", "")
	return result, nil
}

func (lr *LibraryRepository) Update(songID string, updSong models.Song) error {
	debugLogg := func(result, err string) {
		lr.Logs.DebugLog(
			"Update song db method",
			slog.Group("params",
				slog.String("songID", songID),
				slog.Any("starts", updSong),
			),
			slog.String("error", err),
			slog.String("result", result),
		)
	}

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

	_, err := lr.db.Exec(context.Background(), query, updSong.Group, updSong.Song, updSong.ReleaseDate, updSong.Text, updSong.Link, songID)
	if err != nil {
		debugLogg("failure", err.Error())
		return err
	}

	debugLogg("success", "")
	return nil
}

func (lr *LibraryRepository) Delete(songID string) error {
	debugLogg := func(result, err string) {
		lr.Logs.DebugLog(
			"Delete song db method",
			slog.Group("params",
				slog.String("songID", songID),
			),
			slog.String("error", err),
			slog.String("result", result),
		)
	}

	_, err := lr.db.Exec(context.Background(), `DELETE FROM "songs" WHERE id = $1`, songID)
	if err != nil {
		debugLogg("failure", err.Error())
		return err
	}

	debugLogg("success", "")
	return nil
}
