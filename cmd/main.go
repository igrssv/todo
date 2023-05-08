package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/servise"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // driver db
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config: %s", err.Error())
	}
	//init password to env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env: %s", err.Error())
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
		logrus.Fatalf("fail connet db: %s", err.Error())
	}

	// init repository, services and handlers
	repos := repository.NewRepository(db)
	services := servise.NewServive(repos)
	handlers := handler.NewHandler(services)

	//run server
	svr := new(todo.Server)

	go func() {
		if err := svr.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error running http server %s", err.Error())
		}
	}()

	logrus.Print("Start App")

	// add signal for quit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	//shutdown server
	if err := svr.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on sever shutdown: %s", err.Error())
	}

	//disconect db
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}

// init config
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
