package main

import (
	"log"
	"os"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/servise"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // driver db

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}
	//init password to env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env: %s", err.Error())
	}
	// init db
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("fail connet db: %s", err.Error())
	}

	// init repository, services and handlers
	repos := repository.NewRepository(db)
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
