package task

type ITaskRepository interface {
	FindById(id int) (*Task, error)
	FindByTitle(title string) (*Task, error)
	FindAll() ([]*Task, error)
	Save(task *Task) error
	Delete(task *Task) error
}
