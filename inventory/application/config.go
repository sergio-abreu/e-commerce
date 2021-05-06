package application

import "github.com/sergio-vaz-abreu/inventory/infrastructure/postgres"

type Config struct {
	PostgresConfig postgres.Config
	Address        string
}
