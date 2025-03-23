package task

import (
	"fmt"

	"github.com/google/uuid"
	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
	DomainService "github.com/yoshi-d-24/goal-sync/domain/services/task"
)

type UUIDGenerator func() uuid.UUID

var DefaultUUIDGenerator = func() uuid.UUID {
	return uuid.New()
}

type RegisterTaskCommand struct {
	title       string
	description string
	dod         string
	status      int
}

type RegisterTaskApplicationService struct {
	repository    TaskModel.ITaskRepository
	domainService DomainService.ITaskDomainService
	uuidGenerator UUIDGenerator
}

func NewRegisterTaskApplicationService(repository TaskModel.ITaskRepository) *RegisterTaskApplicationService {
	domainService := DomainService.NewTaskDomainService(repository)
	return &RegisterTaskApplicationService{repository: repository, domainService: domainService, uuidGenerator: DefaultUUIDGenerator}
}

func (s *RegisterTaskApplicationService) execute(command RegisterTaskCommand) error {
	task, err := toTaskModel(command, s.uuidGenerator)
	if err != nil {
		return err
	}

	existsDuplicateTitle, err := s.domainService.ExistsDuplicateTitle(task)
	if err != nil {
		return err
	}

	if existsDuplicateTitle {
		return fmt.Errorf("task title is duplicated")
	}

	if err = s.repository.Save(task); err != nil {
		return err
	}

	return nil
}

func toTaskModel(command RegisterTaskCommand, uuidGenerator UUIDGenerator) (*TaskModel.Task, error) {
	id, err := VO.NewTaskId(uuidGenerator().String())
	if err != nil {
		return nil, err
	}

	title, err := VO.NewTitle(command.title)
	if err != nil {
		return nil, err
	}

	description, err := VO.NewTaskDescription(command.description)
	if err != nil {
		return nil, err
	}

	dod, err := VO.NewDoD(command.dod)
	if err != nil {
		return nil, err
	}

	status, err := VO.NewTaskStatus(command.status)
	if err != nil {
		return nil, err
	}

	return TaskModel.NewTask(*id, *title, *description, *status, *dod), nil
}
