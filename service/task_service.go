package service

import (
	"errors"

	"github.com/taufiksty/to-do-list-app-text/entity"
	"github.com/taufiksty/to-do-list-app-text/repository"
)

type TaskService struct {
	Repository repository.TaskRepository
}

func (service TaskService) GetAll() ([]*entity.Task, error) {
	tasks := service.Repository.FindAll()

	if tasks == nil {
		return tasks, errors.New("unfortunately, task not found")
	} else {
		return tasks, nil
	}
}

func (service TaskService) AddTask(task *entity.Task) (*entity.Task, error) {
	newTask := service.Repository.Create(task)

	if newTask == nil {
		return newTask, errors.New("unfortunately, task not created")
	} else {
		return newTask, nil
	}
}

func (service TaskService) UpdateTask(id int, task *entity.Task) (*entity.Task, error) {
	updatedTask := service.Repository.Update(id, task)

	if updatedTask == nil {
		return updatedTask, errors.New("unfortunately, update task fail")
	} else {
		return updatedTask, nil
	}
}

func (service TaskService) DestroyTask(id int) (*entity.Task, error) {
	deletedTask := service.Repository.Delete(id)

	if deletedTask == nil {
		return deletedTask, errors.New("unfortunately, delete task fail")
	} else {
		return deletedTask, nil
	}
}
