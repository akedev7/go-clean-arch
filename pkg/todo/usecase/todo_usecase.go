package usecase

import "github.com/akedev7/go-clean-arch/pkg/interfaces"

type todoUseCase struct {
	todoRepository interfaces.ToDoRepository
}

func (t *todoUseCase) GetByID(id int64) (interfaces.ToDo, error) {
	return t.todoRepository.GetByID(1)
}

func NewTodoUseCase(repo interfaces.ToDoRepository) interfaces.ToDoUsecase {
	return &todoUseCase{
		todoRepository: repo,
	}
}
