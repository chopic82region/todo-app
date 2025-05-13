package main

import (
	"log"

	"github.com/chopic82region/todo-app"
	"github.com/chopic82region/todo-app/pkg/handler"
)

func main(){
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil{
		log.Fatalf("error")
	}
}