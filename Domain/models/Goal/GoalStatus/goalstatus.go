package goal

import (
	"fmt"
)

type GoalStatusEnum int

const (
	Incomplete GoalStatusEnum = iota
	Complete
)

type GoalStatusInterface interface {
	Value() GoalStatusEnum
}

type GoalStatus struct {
	value GoalStatusEnum
}

func NewGoalStatus(value int) (GoalStatusInterface, error) {
	if err := validate(value); err != nil {
		return nil, err
	}
	goalStatus := GoalStatusEnum(value)
	return &GoalStatus{value: goalStatus}, nil
}

func (g *GoalStatus) Value() GoalStatusEnum {
	return g.value
}

func validate(value int) error {
	if value < int(Incomplete) || value > int(Complete) {
		return fmt.Errorf("invalid GoalStatusEnum value: %d", value)
	}

	return nil
}
