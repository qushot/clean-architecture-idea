package dummy

import (
	"fmt"
	"github.com/qushot/clean-arcihtecture-idea/domain/task"
)

type taskRepository struct{}

func (d *taskRepository) Save(entity *task.Entity) error {
	fmt.Printf("タスクを保存するふり(%s)\n", entity.Name)
	return nil
}

func (d *taskRepository) Delete(id task.ID) error {
	fmt.Printf("タスクを削除するふり(%s)\n", id)
	return nil
}

func NewTaskRepository() task.Repository {
	return &taskRepository{}
}
