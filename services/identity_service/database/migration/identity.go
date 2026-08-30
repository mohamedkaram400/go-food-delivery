package migration

import (
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Run(path string, databaseURL string) error {
    m, err := migrate.New(path, databaseURL)

    if err != nil {
        return err
    }

    defer m.Close()

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return err
    }
    return nil
}

