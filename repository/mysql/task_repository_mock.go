package mysql

import (
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"github.com/taufiksty/to-do-list-app-text/entity"
)

// MockTaskRepository is a mock implementation of TaskRepository for testing.
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) FindAll(ctx context.Context) ([]entity.Task, error) {
	args := m.Mock.Called(ctx)
	if args.Get(0) == nil {
		return nil, errors.New("no data")
	} else {
		return args.Get(0).([]entity.Task), nil
	}
}

func (m *MockTaskRepository) Create(ctx context.Context, task entity.Task) (entity.Task, error) {
	args := m.Mock.Called(ctx, task)
	if args.Get(0) == nil {
		return task, errors.New("failed insert data")
	} else {
		return args.Get(0).(entity.Task), nil
	}
}

func (m *MockTaskRepository) Update(ctx context.Context, id int, task entity.Task) (entity.Task, error) {
	args := m.Mock.Called(ctx, id, task)
	if args.Get(0) == nil {
		return task, errors.New("failed update data")
	} else {
		return args.Get(0).(entity.Task), nil
	}
}

func (m *MockTaskRepository) Delete(ctx context.Context, id int) error {
	args := m.Mock.Called(ctx, id)
	if args.Get(0) == nil {
		return errors.New("failed delete data")
	}
	return nil
}
