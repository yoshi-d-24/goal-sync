package task

type ITaskRepository interface {
	FindById(id string) (*Task, error)
	FindByTitle(title string) (*Task, error)
	FindAll() ([]*Task, error)
	Save(task *Task) error
	Delete(id int) error
}
