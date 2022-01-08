package main

import (
	"fmt"
	"os"

	"github.com/akedev7/go-clean-arch/pkg/todo/delivery/http"
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
	router := gin.New()
	usecase := usecase.NewTodoUseCase()
	http.NewToDoHandler(router, usecase)
	router.Run()
	return nil
}
