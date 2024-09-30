package repository

import (
	"MusicLibrary/models"
	"MusicLibrary/pkg/logger"
	"context"

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

func (lc *LibraryRepository) GetSongText() {

}

func (lc *LibraryRepository) Update() {

}

func (lc *LibraryRepository) Delete(songID string) error {
	_, err := lc.db.Exec(context.Background(), `DELETE FROM "songs" WHERE id = $1`, songID)
	if err != nil {
		return err
	}

	return nil
}
