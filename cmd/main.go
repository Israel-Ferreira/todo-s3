package main

import (
	"log"
	"net/http"

	"github.com/Israel-Ferreira/todo-s3/internal/config"
	"github.com/Israel-Ferreira/todo-s3/internal/controllers"
	"github.com/Israel-Ferreira/todo-s3/internal/routes"
)

func init() {
	config.LoadEnvVars()
}

func main() {

	router := routes.CreateUserRouter(&controllers.UserControllerImpl{})

	log.Println("Servidor Iniciado na porta: ", config.ConfigVars.Port)
	log.Fatalln(http.ListenAndServe(config.ConfigVars.Port, router))
}
