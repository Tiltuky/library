package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"library/data"
	"library/internal/controller"
	"library/internal/repository"
	"library/internal/service"
	"library/server"
	"os"
	"os/signal"
	"syscall"

	_ "library/docs" // import generated docs
)

// @title Library API
// @version 1.0
// @description This is a sample server Library server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// init DB
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	// generate fake data
	if err := data.InitData(db); err != nil {
		logrus.Fatalf("failed to initialize data: %v", err)
	}

	// init repositories
	repos := repository.NewRepository(db)
	// init service
	services := service.NewService(repos)
	// init controller
	handlers := controller.NewHandler(services)

	// run http server
	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Library Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Library Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
