//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/qushot/clean-arcihtecture-idea/application/usecase/addtask"
	"github.com/qushot/clean-arcihtecture-idea/infrastructure/persistence/dummy"
	"github.com/qushot/clean-arcihtecture-idea/interfaces/controller"
	"github.com/qushot/clean-arcihtecture-idea/interfaces/presenter"
)

func InitializeApp() *controller.AddTask {
	wire.Build(
		controller.NewAddTask,
		addtask.NewInteractor,
		presenter.NewAddTask,
		dummy.NewTaskRepository,
	)
	return &controller.AddTask{}
}
