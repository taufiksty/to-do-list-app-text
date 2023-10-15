package im

import (
	"github.com/stretchr/testify/mock"
	"github.com/taufiksty/to-do-list-app-text/entity"
)

type TaskRepositoryMock struct {
	Mock mock.Mock
}

func (repository *TaskRepositoryMock) FindAll() []*entity.Task {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	} else {
		return arguments.Get(0).([]*entity.Task)
	}
}

func (repository *TaskRepositoryMock) Create(task *entity.Task) *entity.Task {
	arguments := repository.Mock.Called(task)
	if arguments.Get(0) == nil {
		return nil
	} else {
		return arguments.Get(0).(*entity.Task)
	}
}

func (repository *TaskRepositoryMock) Update(id int, task *entity.Task) *entity.Task {
	arguments := repository.Mock.Called(id, task)
	if arguments.Get(0) == nil {
		return nil
	} else {
		return arguments.Get(0).(*entity.Task)
	}
}

func (repository *TaskRepositoryMock) Delete(id int) *entity.Task {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		return arguments.Get(0).(*entity.Task)
	}
}
