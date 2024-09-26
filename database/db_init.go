package database

import (
	"MusicLibrary/configs"
	"context"

	"github.com/jackc/pgx/v4"
)

func DBInit(cfg configs.DBConfig) (*pgx.Conn, error) {
	dbUrl := "postgres://" + cfg.DbUsername + ":" + cfg.DbPassword + "@" + cfg.DbHost + ":" + cfg.DbPort + "/" + cfg.DbName
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
