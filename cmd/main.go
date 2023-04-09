package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/servise"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	// init repository, services and handlers
	repos := repository.NewRepository()
	services := servise.NewServive(repos)
	handlers := handler.NewHandler(services)

	//run server
	svr := new(todo.Server)
	if err := svr.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server %s", err.Error())
	}

}

// init config
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
