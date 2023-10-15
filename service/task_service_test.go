package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/taufiksty/to-do-list-app-text/entity"
	"github.com/taufiksty/to-do-list-app-text/repository/mysql"
)

var taskRepository = &mysql.MockTaskRepository{Mock: mock.Mock{}}
var taskService = TaskService{Repository: taskRepository}

var layout = "2006-01-02 15:04:05 -0700 MST"

/*
* Unit Test
 */
func TestTaskService_FindAll(t *testing.T) {
	// arrange
	parsedTime, _ := time.Parse(layout, "2023-10-15 07:41:02 +0000 UTC")
	expectedTasks := []entity.Task{
		{
			Id:          1,
			Title:       "Task 1",
			Description: "Description 1",
			Done:        false,
			CreatedAt:   parsedTime,
			UpdatedAt:   parsedTime,
		},
		{
			Id:          2,
			Title:       "Task 2",
			Description: "Description 2",
			Done:        true,
			CreatedAt:   parsedTime,
			UpdatedAt:   parsedTime,
		},
	}

	ctx := context.Background()

	// process mock
	taskRepository.Mock.On("FindAll", ctx).Return(expectedTasks)

	// action
	result, err := taskService.GetAll()

	// assert
	taskRepository.Mock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2, "Result must have length 2")
	assert.Equal(t, expectedTasks, result, "Result is not same as expected")
}

func TestTaskService_AddTask(t *testing.T) {
	// arrange
	parsedTime, _ := time.Parse(layout, "2023-10-15 07:41:02 +0000 UTC")
	newTask := entity.Task{
		Title:       "New Task",
		Description: "This is a new task",
	}
	addedTask := entity.Task{
		Id:          3,
		Title:       "New Task",
		Description: "This is a new task",
		Done:        false,
		CreatedAt:   parsedTime,
		UpdatedAt:   parsedTime,
	}

	ctx := context.Background()

	// process mock
	taskRepository.Mock.On("Create", ctx, newTask).Return(addedTask)

	// action
	result, err := taskService.AddTask(newTask)

	// assert
	taskRepository.Mock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, addedTask, result, "Result is not same as expected")
}

func TestTaskService_UpdateTask(t *testing.T) {
	// arrange
	parsedTime, _ := time.Parse(layout, "2023-10-15 07:41:02 +0000 UTC")
	updateTask := entity.Task{
		Title:       "Update Task",
		Description: "This is an updated task",
		Done:        true,
	}
	updatedTask := entity.Task{
		Id:          3,
		Title:       "Update Task",
		Description: "This is an updated task",
		Done:        true,
		CreatedAt:   parsedTime,
		UpdatedAt:   parsedTime,
	}
	indexUpdatedTask := 3

	ctx := context.Background()

	// process mock
	taskRepository.Mock.On("Update", ctx, indexUpdatedTask, updateTask).Return(updatedTask)

	// action
	result, err := taskService.UpdateTask(indexUpdatedTask, updateTask)

	// assert
	taskRepository.Mock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedTask, result, "Result is not same as expected")
}

func TestTaskService_DestroyTask(t *testing.T) {
	// arrange
	indexDeletedTask := 3
	ctx := context.Background()

	// process mock
	taskRepository.Mock.On("Delete", ctx, indexDeletedTask).Return(nil)

	// action
	result, err := taskService.DestroyTask(indexDeletedTask)

	// assert
	taskRepository.Mock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, true, result, "Result is not same as expected")
}

/*
- Benchmark
*/
func BenchmarkTaskService(t *testing.B) {
	t.Run("GetAll", func(b *testing.B) {
		parsedTime, _ := time.Parse(layout, "2023-10-15 07:41:02 +0000 UTC")
		expectedTasks := []entity.Task{
			{
				Id:          1,
				Title:       "Task 1",
				Description: "Description 1",
				Done:        false,
				CreatedAt:   parsedTime,
				UpdatedAt:   parsedTime,
			},
			{
				Id:          2,
				Title:       "Task 2",
				Description: "Description 2",
				Done:        true,
				CreatedAt:   parsedTime,
				UpdatedAt:   parsedTime,
			},
		}

		ctx := context.Background()

		taskRepository.Mock.On("FindAll", ctx).Return(expectedTasks)

		for i := 0; i < b.N; i++ {
			_, _ = taskService.GetAll()
		}
	})

	t.Run("Create", func(b *testing.B) {
		parsedTime, _ := time.Parse(layout, "2023-10-15 07:41:02 +0000 UTC")
		newTask := entity.Task{
			Title:       "New Task",
			Description: "This is a new task",
		}
		addedTask := entity.Task{
			Id:          3,
			Title:       "New Task",
			Description: "This is a new task",
			Done:        false,
			CreatedAt:   parsedTime,
			UpdatedAt:   parsedTime,
		}

		ctx := context.Background()

		taskRepository.Mock.On("Create", ctx, newTask).Return(addedTask)

		for i := 0; i < b.N; i++ {
			_, _ = taskService.AddTask(newTask)
		}
	})

	t.Run("Update", func(b *testing.B) {
		parsedTime, _ := time.Parse(layout, "2023-10-15 07:41:02 +0000 UTC")
		updateTask := entity.Task{
			Title:       "Update Task",
			Description: "This is an updated task",
			Done:        true,
		}
		updatedTask := entity.Task{
			Id:          1,
			Title:       "Update Task",
			Description: "This is an updated task",
			Done:        true,
			CreatedAt:   parsedTime,
			UpdatedAt:   parsedTime,
		}
		indexUpdatedTask := 1

		ctx := context.Background()

		// process mock
		taskRepository.Mock.On("Update", ctx, indexUpdatedTask, updateTask).Return(updatedTask)

		for i := 0; i < b.N; i++ {
			_, _ = taskService.UpdateTask(indexUpdatedTask, updateTask)
		}
	})

	t.Run("Delete", func(b *testing.B) {
		indexDeletedTask := 1
		ctx := context.Background()

		// process mock
		taskRepository.Mock.On("Delete", ctx, indexDeletedTask).Return(nil)

		for i := 0; i < b.N; i++ {
			_, _ = taskService.DestroyTask(indexDeletedTask)
		}
	})
}
