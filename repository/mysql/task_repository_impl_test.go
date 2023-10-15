package mysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/taufiksty/to-do-list-app-text/database"
	"github.com/taufiksty/to-do-list-app-text/entity"
)

func TestTaskFindAll(t *testing.T) {
	taskRepository := NewTaskRepository(database.GetConnection())

	ctx := context.Background()

	result, err := taskRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, task := range result {
		fmt.Println(task)
	}
}

func TestTaskCreate(t *testing.T) {
	taskRepository := NewTaskRepository(database.GetConnection())

	ctx := context.Background()

	task := entity.Task{
		Title:       "New Task",
		Description: "Description of new task",
	}

	result, err := taskRepository.Create(ctx, task)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestTaskUpdate(t *testing.T) {
	taskRepository := NewTaskRepository(database.GetConnection())

	ctx := context.Background()

	id := 6
	task := entity.Task{
		Title:       "Update task",
		Description: "Update description task",
		Done:        true,
	}

	result, err := taskRepository.Update(ctx, id, task)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestTaskDelete(t *testing.T) {
	taskRepository := NewTaskRepository(database.GetConnection())

	ctx := context.Background()

	id := 6

	err := taskRepository.Delete(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success deleted")
}
