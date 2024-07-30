package app

import "kirmac-site-backend/common/postgresql"

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgreSqlConfig := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostgreSqlConfig: postgreSqlConfig,
	}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "5433",
		UserName:              "kirmac",
		Password:              "kirmac123",
		DbName:                "kirmac_site",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
