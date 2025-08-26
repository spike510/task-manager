package task

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(userID int, title, description string) (*Task, error) {
	var t Task
	err := r.db.QueryRow(
		"INSERT INTO tasks (user_id, title, description) VALUES ($1, $2, $3) RETURNING id, user_id, title, description, done, created_at",
		userID, title, description,
	).Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Done, &t.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *Repository) GetAllByUser(userID int) ([]Task, error) {
	rows, err := r.db.Query("SELECT id, user_id, title, description, done, created_at FROM tasks WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Done, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *Repository) Update(id int, userID int, title, description string, done bool) (*Task, error) {
	var t Task
	err := r.db.QueryRow(
		"UPDATE tasks SET title=$1, description=$2, done=$3 WHERE id=$4 AND user_id=$5 RETURNING id, user_id, title, description, done, created_at",
		title, description, done, id, userID,
	).Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Done, &t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repository) Delete(id int, userID int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id=$1 AND user_id=$2", id, userID)
	return err
}
