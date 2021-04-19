package addtask

import "github.com/qushot/clean-arcihtecture-idea/domain/task"

type OutputPort interface {
	// Do TODO めっちゃ適当にDoって名前つけた
	Do(entity *task.Entity) error
}
