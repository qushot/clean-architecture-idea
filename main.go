package main

import (
	"fmt"
	"github.com/qushot/clean-arcihtecture-idea/domain/task"
	"log"
)

func main() {
	ctrl := InitializeApp()

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
