package im

import "github.com/taufiksty/to-do-list-app-text/entity"

type TaskRepository interface {
	FindAll() []*entity.Task
	Create(task *entity.Task) *entity.Task
	Update(id int, task *entity.Task) *entity.Task
	Delete(id int) *entity.Task
}
