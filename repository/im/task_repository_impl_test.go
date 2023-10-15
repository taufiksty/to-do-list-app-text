package im

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taufiksty/to-do-list-app-text/entity"
)

/*
* Unit Test
 */
func TestInMemoryTaskRepository(t *testing.T) {
	repo := InitInMemoryTaskRepository()

	t.Run("FindAll", func(t *testing.T) {
		// action
		tasks := repo.FindAll()

		// assert
		assert.Len(t, tasks, 2, "Expected 2 tasks")
	})

	t.Run("Create", func(t *testing.T) {
		// arrange
		newTask := &entity.Task{
			Title:       "New Task",
			Description: "This is a new task",
			Done:        false,
		}

		// action & assert
		createdTask := repo.Create(newTask)

		assert.NotNil(t, createdTask, "Expected created task to not be nil")

		tasks := repo.FindAll()

		assert.Len(t, tasks, 3, "Expected 3 tasks after creating new task")
	})

	t.Run("Update", func(t *testing.T) {
		// arrange
		updatedTask := &entity.Task{
			Title:       "Updated Task",
			Description: "This is an updated task",
			Done:        true,
		}

		updatedTaskIndex := 2

		oldTask := *repo.find(updatedTaskIndex + 1)

		// action
		updated := repo.Update(updatedTaskIndex+1, updatedTask)

		// assert
		assert.NotNil(t, updated, "Expected updated to not be nil")
		tasks := repo.FindAll()
		updatedTaskInRepo := tasks[updatedTaskIndex]
		assert.Equal(t, updatedTaskInRepo, updatedTask, "Task in the repository was not update correctly")

		assert.False(t, repo.taskExists(oldTask.Title), "Old task must be updated")
	})

	t.Run("Delete", func(t *testing.T) {
		// assert
		deletedTaskIndex := 2
		deletedTask := *repo.find(deletedTaskIndex + 1)

		// action
		deleted := repo.Delete(deletedTaskIndex + 1)

		// assert
		assert.NotNil(t, deleted, "Expected deleted to not be nil")

		tasks := repo.FindAll()

		assert.Len(t, tasks, 2, "Expected 2 tasks after deleted last one")

		assert.False(t, repo.taskExists(deletedTask.Title), "Deleted task should not be exist in the repository")
	})
}

/*
* Benchmark
 */
func BenchmarkInMemoryTaskRepository(t *testing.B) {
	repo := InitInMemoryTaskRepository()

	t.Run("FindAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = repo.FindAll()
		}
	})

	t.Run("Create", func(b *testing.B) {
		newTask := &entity.Task{
			Title:       "New Task",
			Description: "This is a new task",
			Done:        false,
		}

		for i := 0; i < b.N; i++ {
			_ = repo.Create(newTask)
		}
	})

	t.Run("Update", func(b *testing.B) {
		updatedTask := &entity.Task{
			Title:       "Updated Task",
			Description: "This is an updated task",
			Done:        true,
		}
		updatedTaskIndex := 1

		for i := 0; i < b.N; i++ {
			_ = repo.Update(updatedTaskIndex+1, updatedTask)
		}
	})

	t.Run("Delete", func(b *testing.B) {
		deletedTaskIndex := 1

		for i := 0; i < b.N; i++ {
			_ = repo.Delete(deletedTaskIndex + 1)
		}
	})
}
