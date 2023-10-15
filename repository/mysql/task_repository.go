package mysql

import (
	"context"

	"github.com/taufiksty/to-do-list-app-text/entity"
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]entity.Task, error)
	Create(ctx context.Context, task entity.Task) (entity.Task, error)
	Update(ctx context.Context, id int, task entity.Task) (entity.Task, error)
	Delete(ctx context.Context, id int) error
}
