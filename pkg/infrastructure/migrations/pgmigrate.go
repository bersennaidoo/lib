package migrations

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func PGMigrate(db *sql.DB) {

	migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Errorf("%w", err)
	}

	migrator, err := migrate.NewWithDatabaseInstance("file:///home/bersen/go_projects/filmdb/pkg/migrations", "postgres", migrationDriver)
	if err != nil {
		fmt.Errorf("%w", err)
	}

	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Errorf("%w", err)
	}
}
