package repository

import (
	"database/sql"
	"fmt"

	"github.com/golang-web/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return TaskRepository{db: db}
}

func (r *TaskRepository) Create(payload *model.Task) error {
	query := `INSERT INTO task(name, status) VALUES($1, $2)`
	_, err := r.db.Exec(query, payload.Name, payload.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) GetAllTask() ([]*model.Task, error) {
	query := `SELECT id, name, status FROM task`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var tasks []*model.Task

	for rows.Next() {
		task := &model.Task{}
		if err := rows.Scan(&task.ID, &task.Name, &task.Status); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return tasks, nil
}
