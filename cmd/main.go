package main

import (
	"CourseProject"
	"CourseProject/pkg/handler"
	"CourseProject/pkg/repository"
	"CourseProject/pkg/service"
	"github.com/spf13/viper"

	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(courseProject.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRouters()); err != nil {
		log.Fatalf("error ocured while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
