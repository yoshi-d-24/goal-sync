package task

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid"`
	Title       string
	Description string
	Dod         string
	Status      uint8
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type GormTaskRepository struct {
	db *gorm.DB
	TaskModel.ITaskRepository
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository {
	return &GormTaskRepository{db: db}
}

func (r *GormTaskRepository) FindById(id string) (*TaskModel.Task, error) {
	var task Task

	u, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}
	err = r.db.First(&task, u).Error

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
	var tasks []Task
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
	var task Task
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
	taskRecord, err := toTaskRecord(task)

	if err != nil {
		return err
	}

	err = r.db.Create(&taskRecord).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *GormTaskRepository) Delete(id string) error {
	u, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	err = r.db.Delete(&Task{}, u).Error

	if err != nil {
		return err
	}

	return nil
}

func toTaskRecord(taskModel *TaskModel.Task) (*Task, error) {
	id, err := uuid.Parse(taskModel.Id().Value())

	if err != nil {
		return nil, err
	}
	return &Task{
		ID:          id,
		Title:       taskModel.Title().Value(),
		Description: taskModel.Description().Value(),
		Dod:         taskModel.DoD().Value(),
		Status:      uint8(taskModel.Status().Value()),
	}, nil
}

func toTaskModel(task *Task) (*TaskModel.Task, error) {
	taskId, err := VO.NewTaskId(task.ID.String())
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
	status, err := VO.NewTaskStatus(int(task.Status))
	if err != nil {
		return nil, err
	}

	return TaskModel.NewTask(*taskId, *title, *description, *status, *dod), nil
}
