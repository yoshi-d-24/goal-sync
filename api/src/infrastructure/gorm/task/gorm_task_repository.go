package task

import (
	"errors"

	"gorm.io/gorm"

	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
)

type TaskDto struct {
	gorm.Model
	Title       string
	Description string
	Dod         string
	Status      int
}

type GormTaskRepository struct {
	db *gorm.DB
	TaskModel.ITaskRepository
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository {
	return &GormTaskRepository{db: db}
}

func (r *GormTaskRepository) FindById(id int) (*TaskModel.Task, error) {
	var task TaskDto
	err := r.db.First(&task, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return toTaskModel(&task)
}

func (r *GormTaskRepository) FindAll() ([]*TaskModel.Task, error) {
	var tasks []TaskDto
	err := r.db.Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	var taskModels []*TaskModel.Task
	for _, task := range tasks {
		taskModel, err := toTaskModel(&task)
		if err != nil {
			return nil, err
		}
		taskModels = append(taskModels, taskModel)
	}

	return taskModels, nil
}

func (r *GormTaskRepository) FindByTitle(title string) (*TaskModel.Task, error) {
	var task TaskDto
	err := r.db.Where("title = ?", title).First(&task).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return toTaskModel(&task)
}

func (r *GormTaskRepository) Save(task *TaskModel.Task) error {
	taskDto := toTaskDto(task)
	err := r.db.Create(&taskDto).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *GormTaskRepository) Delete(id int) error {
	err := r.db.Delete(&TaskDto{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func toTaskDto(taskModel *TaskModel.Task) *TaskDto {
	return &TaskDto{
		Title:       taskModel.Title().Value(),
		Description: taskModel.Description().Value(),
		Dod:         taskModel.DoD().Value(),
		Status:      taskModel.Status().Value(),
	}
}

func toTaskModel(task *TaskDto) (*TaskModel.Task, error) {
	taskId, err := VO.NewTaskId(int(task.ID))
	if err != nil {
		return nil, err
	}
	title, err := VO.NewTitle(task.Title)
	if err != nil {
		return nil, err
	}
	description, err := VO.NewTaskDescription(task.Description)
	if err != nil {
		return nil, err
	}
	dod, err := VO.NewDoD(task.Dod)
	if err != nil {
		return nil, err
	}
	status, err := VO.NewTaskStatus(task.Status)
	if err != nil {
		return nil, err
	}

	return TaskModel.NewTask(*taskId, *title, *description, *status, *dod), nil
}
