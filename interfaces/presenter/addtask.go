package presenter

import (
	"fmt"
	"github.com/qushot/clean-arcihtecture-idea/application/usecase/addtask"
	"github.com/qushot/clean-arcihtecture-idea/domain/task"
)

type addTask struct{}

func NewAddTask() addtask.OutputPort {
	return &addTask{}
}

func (p *addTask) Do(entity *task.Entity) error {
	fmt.Println(entity)
	return nil
}
