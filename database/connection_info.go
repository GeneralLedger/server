package database

import (
	"fmt"
)

type ConnectionInfo struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
	SSLMode      string
}

func (c ConnectionInfo) ToURI() string {
	result := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DatabaseName,
		c.SSLMode,
	)
	return result
}
