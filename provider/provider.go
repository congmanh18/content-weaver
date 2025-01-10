package provider

import (
	conf "github.com/congmanh18/content-weaver/configs"

	"github.com/congmanh18/content-weaver/core/db/postgresql"
)

type AppProvider struct {
	Postgres *postgresql.Database
	// Redis    *redis.Client
}

func NewAppProvider(config conf.ServiceConfig) *AppProvider {
	return &AppProvider{
		// Redis:    ProvideRedis(config),
		Postgres: ProvidePostgres(config),
	}
}
