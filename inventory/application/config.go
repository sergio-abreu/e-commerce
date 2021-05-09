package application

import (
	"github.com/Netflix/go-env"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/inventory/infrastructure/postgres"
)

func LoadConfigFromEnv() (Config, error) {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	return config, errors.Wrap(err, "failed to get config from environment")
}

type Config struct {
	PostgresConfig postgres.Config
	RpcAddress     string `env:"RPC_ADDRESS"`
}
