package main

import (
	"context"
	"errors"
	"ipdb/bootstrap"
	"ipdb/modules/config"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type invokeDeps struct {
	fx.In

	Rest *bootstrap.Serve
}

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config.yaml file not found")
		}
	}

	err := viper.Unmarshal(&config.GlobalConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Println(config.GlobalConfig.IPIP.Path)
	log.Println(config.GlobalConfig.IP2Location.Path)

	// telemetrymod.FxOptions(),
	// sensitivemod.FxOptions(),
	app := fx.New(
		bootstrap.FxOptions(),
		fx.Invoke(func(lc fx.Lifecycle, deps invokeDeps) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go deps.Rest.Start()
					log.Println("Press Ctrl+C to exit")
					return nil
				},
				OnStop: func(ctx context.Context) error {
					//deps.LokiWriter.Shutdown()
					return errors.Join(deps.Rest.Stop())
				},
			})
		}),
	)

	app.Run()
}
