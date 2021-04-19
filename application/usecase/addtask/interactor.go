package addtask

import "github.com/qushot/clean-arcihtecture-idea/domain/task"

type interactor struct {
	taskRepository task.Repository
	output         OutputPort
}

func NewInteractor(taskRepository task.Repository, output OutputPort) InputPort {
	return &interactor{
		taskRepository: taskRepository,
		output:         output,
	}
}

func (r *interactor) Handle(entity *task.Entity) error {
	if err := r.taskRepository.Save(entity); err != nil {
		return err
	}

	return r.output.Do(entity)
}
