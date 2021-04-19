package main

import (
	"fmt"
	"github.com/qushot/clean-arcihtecture-idea/application/usecase/addtask"
	"github.com/qushot/clean-arcihtecture-idea/domain/task"
	"github.com/qushot/clean-arcihtecture-idea/infrastructure/persistence/dummy"
	"github.com/qushot/clean-arcihtecture-idea/interfaces/controller"
	"github.com/qushot/clean-arcihtecture-idea/interfaces/presenter"
	"log"
)

func main() {
	repo := dummy.NewTaskRepository()
	output := presenter.NewAddTask()
	input := addtask.NewInteractor(repo, output)
	ctrl := controller.NewAddTask(input)

	var n, d string
	fmt.Println("Input to task info")
	fmt.Print("name > ")
	fmt.Scan(&n)
	fmt.Print("description > ")
	fmt.Scan(&d)

	t, err := task.NewEntity(task.Name(n), task.Description(d))
	if err != nil {
		log.Fatalf("失敗: %v", err)
	}

	if err := ctrl.Execute(t); err != nil {
		log.Fatalf("失敗: %v", err)
	}

	fmt.Println("DONE")
}
