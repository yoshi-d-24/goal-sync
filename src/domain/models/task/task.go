package task

import (
	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/valueobject"
)

type TaskInterface interface {
	Id() VO.TaskId
	Title() VO.Title
	Description() VO.TaskDescription
	Status() VO.TaskStatus
	DoD() VO.DoD
	Comments() []VO.TaskComment
}

type Task struct {
	id          VO.TaskId
	title       VO.Title
	description VO.TaskDescription
	status      VO.TaskStatus
	dod         VO.DoD
	comments    []VO.TaskComment
}

func NewTask(id VO.TaskId, title VO.Title, description VO.TaskDescription, status VO.TaskStatus, dod VO.DoD, comments []VO.TaskComment) TaskInterface {
	return &Task{id, title, description, status, dod, comments}
}

func (t *Task) Id() VO.TaskId {
	return t.id
}

func (t *Task) Title() VO.Title {
	return t.title
}

func (t *Task) Description() VO.TaskDescription {
	return t.description
}

func (t *Task) Status() VO.TaskStatus {
	return t.status
}

func (t *Task) DoD() VO.DoD {
	return t.dod
}

func (t *Task) Comments() []VO.TaskComment {
	return t.comments
}

func (t *Task) AddComment(comment VO.TaskComment) {
	t.comments = append(t.comments, comment)
}
