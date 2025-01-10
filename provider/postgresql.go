package provider

import (
	"fmt"

	conf "github.com/congmanh18/content-weaver/configs"
	"github.com/congmanh18/content-weaver/core/db/postgresql"
)

func ProvidePostgres(config conf.ServiceConfig) *postgresql.Database {
	fmt.Println("Connecting to PostgreSQL...")
	return postgresql.New(postgresql.Connection{
		Host:     config.DBHost,
		Port:     config.DBPort,
		Database: config.DBName,
		User:     config.DBUser,
		Password: config.DBPwd,
		SSLMode:  postgresql.Disable,
	})
}

// Disable - Require
