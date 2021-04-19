package main

import (
	"fmt"
	"github.com/qushot/clean-arcihtecture-idea/application/usecase/addtask"
	"github.com/qushot/clean-arcihtecture-idea/domain/task"
	"github.com/qushot/clean-arcihtecture-idea/interfaces/controller"
	"github.com/qushot/clean-arcihtecture-idea/interfaces/presenter"
	"log"
)

type dummyRepo struct{}

func (d *dummyRepo) Save(entity *task.Entity) error {
	fmt.Printf("タスクを保存(%s)\n", entity.Name)
	return nil
}

func (d *dummyRepo) Delete(id task.ID) error {
	fmt.Printf("タスクを削除(%s)\n", id)
	return nil
}

func NewDummyRepo() task.Repository {
	return &dummyRepo{}
}

func main() {
	repo := NewDummyRepo()
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
