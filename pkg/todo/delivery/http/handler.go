package http

import (
	"net/http"

	"github.com/akedev7/go-clean-arch/pkg/interfaces"
	"github.com/gin-gonic/gin"
)

type ToDoHandler struct {
	toDoUseCase interfaces.ToDoUsecase
}

func NewToDoHandler(router *gin.Engine, usecase interfaces.ToDoUsecase) {
	handler := &ToDoHandler{
		toDoUseCase: usecase,
	}

	router.GET("/test", handler.testHandler)
}

func (t *ToDoHandler) testHandler(c *gin.Context) {
	data, _ := t.toDoUseCase.GetByID(1)
	c.JSON(http.StatusOK, data)
}
