package application

import "github.com/sergio-vaz-abreu/discount/infrastructure/postgres"

type Config struct {
	PostgresConfig postgres.Config
	RpcAddress     string `env:"RPC_ADDRESS"`
}
