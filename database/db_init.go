package database

import (
	"MusicLibrary/configs"
	"context"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

func DBInit(cfg configs.DBConfig) (*pgx.Conn, error) {
	dbUrl := "postgres://" + cfg.DbUsername + ":" + cfg.DbPassword + "@" + cfg.DbHost + ":" + cfg.DbPort + "/" + cfg.DbName
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func RunMigrations(db *pgx.Conn) {
	stdDB := stdlib.OpenDB(*db.Config())

	driver, err := postgres.WithInstance(stdDB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create Postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("Could not start migration: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v", err)
	}

	log.Println("Migrations ran successfully")
}
