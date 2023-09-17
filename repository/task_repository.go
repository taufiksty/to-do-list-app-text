package repository

import "github.com/taufiksty/to-do-list-app-text/entity"

type TaskRepository interface {
	FindAll() []*entity.Task
	Create(task *entity.Task) *entity.Task
	Update(id int, task *entity.Task) *entity.Task
	Delete(id int) *entity.Task
}

type InMemoryTaskRepository struct {
	tasks []*entity.Task
}

func InitInMemoryTaskRepository() *InMemoryTaskRepository {
	initialTasks := []*entity.Task{
		{
			Title:       "Belajar",
			Description: "Belajar golang dasar",
			Done:        false,
		},
		{
			Title:       "Olahraga",
			Description: "Jogging santai",
			Done:        true,
		},
	}

	return &InMemoryTaskRepository{tasks: initialTasks}
}

func (r *InMemoryTaskRepository) taskExists(taskTitle string) bool {
	for _, task := range r.tasks {
		if task.Title == taskTitle {
			return true
		}
	}

	return false
}

func (r *InMemoryTaskRepository) FindAll() []*entity.Task {
	if len(r.tasks) == 0 {
		return nil
	}

	return r.tasks
}

func (r *InMemoryTaskRepository) find(id int) *entity.Task {
	return r.tasks[id-1]
}

func (r *InMemoryTaskRepository) Create(task *entity.Task) *entity.Task {
	r.tasks = append(r.tasks, task)

	if r.taskExists(task.Title) {
		return task
	} else {
		return nil
	}
}

func (r *InMemoryTaskRepository) Update(id int, task *entity.Task) *entity.Task {
	r.tasks[id-1].Title = task.Title
	r.tasks[id-1].Description = task.Description
	r.tasks[id-1].Done = task.Done

	if r.taskExists(task.Title) {
		return task
	} else {
		return nil
	}
}

func (r *InMemoryTaskRepository) Delete(id int) *entity.Task {

	task := r.find(id)

	r.tasks = append(r.tasks[:id-1], r.tasks[id:]...)

	if r.taskExists(task.Title) {
		return nil
	} else {
		return task
	}
}
