package main

import (
	"log"
	"net/http"

	"github.com/Israel-Ferreira/todo-s3/internal/config"
	"github.com/Israel-Ferreira/todo-s3/internal/controllers"
	"github.com/Israel-Ferreira/todo-s3/internal/routes"
	implService "github.com/Israel-Ferreira/todo-s3/internal/services"
)

func init() {
	config.LoadEnvVars()

	config.CreateAwsSession("sa-east-1", config.ConfigVars.AwsAccessKey, config.ConfigVars.AwsSecretKey)
}

func main() {

	userSrvc := implService.UserServiceImpl{}

	ctr := controllers.NewUserController(userSrvc)

	router := routes.CreateUserRouter(ctr)

	log.Println("Servidor Iniciado na porta: ", config.ConfigVars.Port)
	log.Fatalln(http.ListenAndServe(config.ConfigVars.Port, router))
}
