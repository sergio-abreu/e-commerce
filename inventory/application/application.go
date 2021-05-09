package application

import (
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/inventory/application/inventory"
	"github.com/sergio-vaz-abreu/inventory/infrastructure/postgres"
	"github.com/sergio-vaz-abreu/inventory/modules/application"
	"github.com/sergio-vaz-abreu/inventory/modules/infrastructure/product"
	"google.golang.org/grpc"
	"net"
)

func Load(config Config) (*Application, error) {
	s := grpc.NewServer()
	db, err := postgres.NewDatabase(config.PostgresConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize postgres database")
	}
	productRepository := product.NewSqlRepository(db)
	inventory.RegisterInventoryServer(s, application.NewInventory(productRepository))
	return &Application{s: s, address: config.RpcAddress}, nil
}

type Application struct {
	s       *grpc.Server
	address string
}

func (a *Application) Run() <-chan error {
	errCh := make(chan error)
	go func() {
		listener, err := net.Listen("tcp", a.address)
		if err != nil {
			errCh <- errors.Wrap(err, "failed to initialize rpc listener")
			return
		}
		if err := a.s.Serve(listener); err != nil {
			errCh <- errors.Wrap(err, "failed while running gRPC server")
			return
		}
	}()
	return errCh
}

func (a *Application) Shutdown() {
	a.s.GracefulStop()
}
