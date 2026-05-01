package students

import (
	"context"
	"database/sql"
	"errors"
	"my-project/internal/models"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetByID(ctx context.Context, id int) (*models.Student, error) {
	query := `SELECT id,name,department_id FROM students WHERE id=$1`

	var s models.Student

	err := r.db.QueryRowContext(ctx, query, id).Scan(&s.ID, &s.Name, &s.DepartmentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}
	return &s, nil
}

func (r *PostgresRepo) Create(ctx context.Context, s *models.Student) error {
	query := `INSERT INTO stundets (name,department_id) VALUES($1,$2)`

	err := r.db.QueryRowContext(ctx, query, s.Name, s.DepartmentID).Scan()

	if err != nil {
		return err
	}

	return nil
}


