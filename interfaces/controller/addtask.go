package controller

import (
	"github.com/qushot/clean-arcihtecture-idea/application/usecase/addtask"
	"github.com/qushot/clean-arcihtecture-idea/domain/task"
)

type addTask struct {
	input addtask.InputPort
}

func NewAddTask(input addtask.InputPort) *addTask {
	return &addTask{
		input: input,
	}
}

func (r *addTask) Execute(entity *task.Entity) error {
	return r.input.Handle(entity)
}
