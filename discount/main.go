package main

import (
	"context"
	env "github.com/Netflix/go-env"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/discount/application"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"os/signal"
)

func main() {
	app := cli.NewApp()
	app.Name = ApplicationName
	app.Description = "An discount application for an ecommerce system"
	app.Version = Version + "(" + GitCommit + ")"
	app.EnableBashCompletion = true
	app.Action = func(cliCtx *cli.Context) error {
		var applicationConfig application.Config
		err := loadConfig(&applicationConfig)
		if err != nil {
			return err
		}
		return run(applicationConfig)
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func run(applicationConfig application.Config) error {
	logrus.Info("starting application")
	app, err := application.Load(applicationConfig)
	if err != nil {
		return err
	}
	appErr := app.Run()
	ctx := gracefullyShutdown()
	defer app.Shutdown()
	select {
	case err := <-appErr:
		return err
	case <-ctx.Done():
		logrus.Info("gracefully shutdown")
		return nil
	}
}

func gracefullyShutdown() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		cancel()
	}()
	return ctx
}

func loadConfig(data interface{}) error {
	_, err := env.UnmarshalFromEnviron(data)
	return errors.Wrap(err, "failed to get config from environment")
}
