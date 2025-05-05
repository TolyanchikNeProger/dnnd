package main

import (
	"log"

	"github.com/TolyanchikNeProger/dnnd"
	"github.com/TolyanchikNeProger/dnnd/pkg/handlers"
	"github.com/TolyanchikNeProger/dnnd/pkg/repository"
	"github.com/TolyanchikNeProger/dnnd/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := new(dnnd.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occired while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
