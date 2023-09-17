package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/taufiksty/to-do-list-app-text/entity"
	"github.com/taufiksty/to-do-list-app-text/repository"
)

var taskRepository = &repository.TaskRepositoryMock{Mock: mock.Mock{}}
var taskService = TaskService{Repository: taskRepository}

/*
* Unit Test
 */
func TestTaskService_GetAll(t *testing.T) {
	// arrange
	expectedTasks := []*entity.Task{
		{
			Title:       "Task 1",
			Description: "Description 1",
			Done:        false,
		},
		{
			Title:       "Task 2",
			Description: "Description 2",
			Done:        true,
		},
	}

	// process mock
	taskRepository.Mock.On("FindAll").Return(expectedTasks)

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
	newTask := &entity.Task{
		Title:       "New Task",
		Description: "This is a new task",
		Done:        false,
	}

	// process mock
	taskRepository.Mock.On("Create", newTask).Return(newTask)

	// action
	result, err := taskService.AddTask(newTask)

	// assert
	taskRepository.Mock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, newTask, result, "Result is not same as expected")
}

func TestTaskService_UpdateTask(t *testing.T) {
	// arrange
	updateTask := &entity.Task{
		Title:       "Update Task",
		Description: "This is an updated task",
		Done:        true,
	}
	indexUpdatedTask := 1

	// process mock
	taskRepository.Mock.On("Update", indexUpdatedTask+1, updateTask).Return(updateTask)

	// action
	result, err := taskService.UpdateTask(indexUpdatedTask+1, updateTask)

	// assert
	taskRepository.Mock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updateTask, result, "Result is not same as expected")
}

func TestTaskService_DestroyTask(t *testing.T) {
	// arrange
	deletedTask := &entity.Task{
		Title:       "Deleted Task",
		Description: "This is an deletedd task",
		Done:        true,
	}
	indexDeletedTask := 1

	// process mock
	taskRepository.Mock.On("Delete", indexDeletedTask+1).Return(deletedTask)

	// action
	result, err := taskService.DestroyTask(indexDeletedTask + 1)

	// assert
	taskRepository.Mock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, deletedTask, result, "Result is not same as expected")
}

/*
* Benchmark
 */
func BenchmarkTaskService(t *testing.B) {
	t.Run("GetAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = taskService.GetAll()
		}
	})

	t.Run("Create", func(b *testing.B) {
		newTask := &entity.Task{
			Title:       "New Task",
			Description: "This is a new task",
			Done:        false,
		}

		for i := 0; i < b.N; i++ {
			_, _ = taskService.AddTask(newTask)
		}
	})

	t.Run("Update", func(b *testing.B) {
		updatedTask := &entity.Task{
			Title:       "Update Task",
			Description: "This is an updated task",
			Done:        true,
		}
		updatedTaskIndex := 1

		for i := 0; i < b.N; i++ {
			_, _ = taskService.UpdateTask(updatedTaskIndex+1, updatedTask)
		}
	})

	t.Run("Delete", func(b *testing.B) {
		deletedTaskIndex := 1

		for i := 0; i < b.N; i++ {
			_, _ = taskService.DestroyTask(deletedTaskIndex + 1)
		}
	})
}
