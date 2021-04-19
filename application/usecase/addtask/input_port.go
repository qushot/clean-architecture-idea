package addtask

import "github.com/qushot/clean-arcihtecture-idea/domain/task"

type InputPort interface {
	Handle(entity *task.Entity) error
}
