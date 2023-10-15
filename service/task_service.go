package service

import (
	"context"
	"errors"

	"github.com/taufiksty/to-do-list-app-text/entity"
	"github.com/taufiksty/to-do-list-app-text/repository/mysql"
)

type TaskService struct {
	Repository mysql.TaskRepository
}

func (service TaskService) GetAll() ([]entity.Task, error) {
	ctx := context.Background()

	tasks, err := service.Repository.FindAll(ctx)

	if err != nil {
		return tasks, errors.New("unfortunately, task not found")
	} else {
		return tasks, nil
	}
}

func (service TaskService) AddTask(task entity.Task) (entity.Task, error) {
	ctx := context.Background()

	newTask, err := service.Repository.Create(ctx, task)
	if err != nil {
		return newTask, errors.New("unfortunately, task not created")
	} else {
		return newTask, nil
	}
}

func (service TaskService) UpdateTask(id int, task entity.Task) (entity.Task, error) {
	ctx := context.Background()

	updatedTask, err := service.Repository.Update(ctx, id, task)

	if err != nil {
		return updatedTask, errors.New("unfortunately, update task fail")
	} else {
		return updatedTask, nil
	}
}

func (service TaskService) DestroyTask(id int) (bool, error) {
	ctx := context.Background()

	err := service.Repository.Delete(ctx, id)

	if err == nil {
		return false, errors.New("unfortunately, delete task fail")
	} else {
		return true, nil
	}
}
