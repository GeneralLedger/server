package database

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEphemeralDatabase(t *testing.T) {
	// Spin up a new Ephemeral Database
	ephemeralDatabase, err := NewEphemeralDatabase(
		ConnectionInfo{
			Host:         os.Getenv("TEST_POSTGRES_HOST"),
			Port:         os.Getenv("TEST_POSTGRES_PORT"),
			User:         os.Getenv("TEST_POSTGRES_USER"),
			Password:     os.Getenv("TEST_POSTGRES_PASSWORD"),
			DatabaseName: os.Getenv("TEST_POSTGRES_DB"),
			SSLMode:      os.Getenv("TEST_POSTGRES_SSL_MODE"),
		},
		Config{
			MigrationsDir: "../database/migrations",
		},
	)
	assert.Nil(t, err)
	defer ephemeralDatabase.Terminate()

	// Test that the database works & we can ping it
	err = ephemeralDatabase.Connection().Ping()
	assert.Nil(t, err)

	// Shut it down
	err = ephemeralDatabase.Terminate()
	assert.Nil(t, err)

	// Verify that we can't ping it
	err = ephemeralDatabase.Connection().Ping()
	assert.NotNil(t, err)
}
