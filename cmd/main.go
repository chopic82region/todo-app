package main

import (
	"log"

	"github.com/chopic82region/todo-app"
	"github.com/chopic82region/todo-app/pkg/handler"
	"github.com/chopic82region/todo-app/pkg/repository"
	"github.com/chopic82region/todo-app/pkg/service"
)

func main(){

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil{
		log.Fatalf("error")
	}
}