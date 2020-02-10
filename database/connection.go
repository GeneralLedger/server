package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Config struct {
	// The Directory Path of the database migrations.  If this is set,
	// the database will be automatically migrated when generating a
	// new connection.
	MigrationsDir string
}

func NewConnection(uri string, config Config) (*sql.DB, error) {
	// Get a connection to the database
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	// If we are passed a MigrationsDir, migrate the database.
	if config.MigrationsDir != "" {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return nil, err
		}
		dbMigrator, err := migrate.NewWithDatabaseInstance(
			"file://"+config.MigrationsDir,
			"postgres",
			driver,
		)
		if err != nil {
			return nil, err
		}
		err = dbMigrator.Up()
		if err != nil {
			if err == migrate.ErrNoChange {
				// Eat the error & log here, we'll consider this a
				// non-error in this context.
				fmt.Println("Migrations up to date.")
			} else {
				return nil, err
			}
		}
	}

	// OK
	return db, nil
}
