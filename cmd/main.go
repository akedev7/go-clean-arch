package main

import (
	"fmt"
	"os"

	"github.com/akedev7/go-clean-arch/pkg/todo/delivery/http"
	"github.com/akedev7/go-clean-arch/pkg/todo/repository"
	"github.com/akedev7/go-clean-arch/pkg/todo/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {

	repo := repository.NewToDoRepository()

	//Inject repo into service
	usecase := usecase.NewTodoUseCase(repo)
	router := gin.New()

	//Inject service into handler
	http.NewToDoHandler(router, usecase)

	router.Run()
	return nil
}
