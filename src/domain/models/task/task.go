package task

import (
	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
)

type Task struct {
	id          VO.TaskId
	title       VO.Title
	description VO.TaskDescription
	status      VO.TaskStatus
	dod         VO.DoD
	// コメントはVOではなくEntityとしてのちのち実装する
	// comments    []VO.TaskComment
}

func NewTask(id VO.TaskId, title VO.Title, description VO.TaskDescription, status VO.TaskStatus, dod VO.DoD) *Task {
	return &Task{id, title, description, status, dod}
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

func (t *Task) Equals(other *Task) bool {
	return t.id.Equals(other.Id()) &&
		t.title.Equals(other.Title()) &&
		t.description.Equals(other.Description()) &&
		t.status.Equals(other.Status()) &&
		t.dod.Equals(other.DoD())
}
