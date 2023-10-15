package mysql

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/taufiksty/to-do-list-app-text/entity"
)

type DBTaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &DBTaskRepository{DB: db}
}

func (repository *DBTaskRepository) FindAll(ctx context.Context) ([]entity.Task, error) {
	query := "SELECT id, title, description, done, created_at, updated_at FROM tasks"

	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task

	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.Id, &task.Title, &task.Description, &task.Done, &task.CreatedAt, &task.UpdatedAt)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repository *DBTaskRepository) Create(ctx context.Context, task entity.Task) (entity.Task, error) {
	query := "INSERT INTO tasks (title, description) VALUES (?, ?)"

	result, err := repository.DB.ExecContext(ctx, query, task.Title, task.Description)
	if err != nil {
		return task, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return task, err
	}

	task.Id = int32(id)
	task.Done = false
	task.CreatedAt = time.Now().UTC().Add(7 * time.Hour)
	task.UpdatedAt = time.Now().UTC().Add(7 * time.Hour)

	return task, err
}

func (repository *DBTaskRepository) Update(ctx context.Context, id int, task entity.Task) (entity.Task, error) {
	query := "UPDATE tasks SET title = ?, description = ?, done = ? WHERE id = ?"

	result, err := repository.DB.ExecContext(ctx, query, task.Title, task.Description, task.Done, id)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return task, err
	}

	if rowsAffected == 0 {
		return task, errors.New("No rows were updated. Task with provided " + strconv.Itoa(id) + " not found.")
	}

	task.Id = int32(id)
	task.UpdatedAt = time.Now().UTC().Add(7 * time.Hour)

	return task, nil

}

func (repository *DBTaskRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM tasks WHERE id = ?"

	result, err := repository.DB.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	if rowsAffected == 0 {
		return errors.New("No rows were updated. Task with provided " + strconv.Itoa(id) + " not found.")
	}

	return nil
}
