package usecase

import "github.com/akedev7/go-clean-arch/pkg/interfaces"

type todoUseCase struct {
}

func (t *todoUseCase) GetByID(id int64) (interfaces.ToDo, error) {
	return interfaces.ToDo{
		ID:    123,
		Title: "Test",
	}, nil
}

func NewTodoUseCase() interfaces.ToDoUsecase {
	return &todoUseCase{}
}
