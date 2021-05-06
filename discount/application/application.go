package application

import (
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/discount/application/discount_calculator"
	"github.com/sergio-vaz-abreu/discount/infrastructure/postgres"
	"github.com/sergio-vaz-abreu/discount/modules/application/calculator"
	"github.com/sergio-vaz-abreu/discount/modules/infrastructure/product"
	"github.com/sergio-vaz-abreu/discount/modules/infrastructure/user"
	"google.golang.org/grpc"
	"net"
)

func Load(config Config) (*Application, error) {
	s := grpc.NewServer()
	db, err := postgres.NewDatabase(config.PostgresConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize postgres database")
	}
	userRepository := user.NewSqlRepository(db)
	productRepository := product.NewSqlRepository(db)
	discount_calculator.RegisterDiscountCalculatorServer(s, calculator.NewCalculator(userRepository, productRepository))
	return &Application{s: s, address: config.Address}, nil
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
