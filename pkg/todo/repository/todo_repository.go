package repository

import "github.com/akedev7/go-clean-arch/pkg/interfaces"

type toDoRepository struct {
}

func (t *toDoRepository) GetByID(id int64) (interfaces.ToDo, error) {
	return interfaces.ToDo{
		ID:    123,
		Title: "Test",
	}, nil
}

func NewToDoRepository() interfaces.ToDoRepository {
	return &toDoRepository{}
}
