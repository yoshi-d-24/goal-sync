package task

// import (
// 	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
// 	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
// 	DomainService "github.com/yoshi-d-24/goal-sync/domain/services/task"
// )

// type RegisterTaskCommand struct {
// 	title       string
// 	description string
// 	dod         string
// 	status      int
// }

// type RegisterTaskApplicationService struct {
// 	repository TaskModel.ITaskRepository
// }

// func NewRegisterTaskApplicationService(repository TaskModel.ITaskRepository) *RegisterTaskApplicationService {
// 	return &RegisterTaskApplicationService{repository: repository}
// }

// func (s *RegisterTaskApplicationService) execute(command RegisterTaskCommand) error {
// 	domainService := DomainService.NewTaskDomainService(s.repository)

// 	task := TaskModel.NewTask()

// 	existsDuplicateTitl := domainService.ExistsDuplicateTitle()
// }

// func toTaskModel(command RegisterTaskCommand) (*TaskModel.Task, error) {
// 	title, err := VO.NewTitle(command.title)
// 	if err != nil {
// 		return nil, err
// 	}
// 	description, err := VO.NewTaskDescription(command.description)
// 	if err != nil {
// 		return nil, err
// 	}
// 	dod, err := VO.NewDoD(command.dod)
// 	if err != nil {
// 		return nil, err
// 	}
// 	status, err := VO.NewTaskStatus(command.status)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return TaskModel.NewTask(nil, *title, *description, *status, *dod), nil
// }
