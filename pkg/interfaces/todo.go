package interfaces

type ToDo struct {
	ID    int64  `json:"id"`
	Title string `json:"title" validate:"required"`
}

type ToDoUsecase interface {
	GetByID(id int64) (ToDo, error)
}
