package controller

import (
	"github.com/qushot/clean-arcihtecture-idea/application/usecase/addtask"
	"github.com/qushot/clean-arcihtecture-idea/domain/task"
)

type AddTask struct {
	input addtask.InputPort
}

func NewAddTask(input addtask.InputPort) *AddTask {
	return &AddTask{
		input: input,
	}
}

func (r *AddTask) Execute(entity *task.Entity) error {
	return r.input.Handle(entity)
}
