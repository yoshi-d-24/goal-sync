package task

type TaskService struct {
	taskRepository ITaskRepository
}

func (ts *TaskService) ExistsSameTitle(task *Task) (bool, error) {
	sameTitleTask, err := ts.taskRepository.FindByTitle(task.title.Value())
	if err != nil {
		return false, err
	}
	if sameTitleTask != nil {
		return false, err
	}
	return true, nil
}
