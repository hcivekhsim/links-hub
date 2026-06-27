package repository

import (
	"context"
	"database/sql"

	"github.com/hcivekhsim/links-hub/models"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]models.Link, error) {
	links := []models.Link{}

	rows, err := r.db.QueryContext(ctx, "SELECT id, title, url, description, created_at, updated_at FROM  links")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		l := models.Link{}

		err := rows.Scan(&l.ID, &l.Title, &l.URL, &l.Desc, &l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			return nil, err
		}
		links = append(links, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}

func (r *PostgresRepo) AddLink(ctx context.Context, link models.Link) (int, error) {
	query := `INSERT INTO links (title, url, description) 
		VALUES ($1, $2, $3) RETURNING id `

	var id int
	err := r.db.QueryRowContext(ctx, query, link.Title, link.URL, link.Desc).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresRepo) ChangeLink(ctx context.Context, id int, link models.LinkUpdate) error {
	query := `UPDATE links SET 
		title = COALESCE(NULLIF($1,''), title), url = COALESCE(NULLIF($2,''), url), description = COALESCE(NULLIF($3,''), description), 
		updated_at = NOW() WHERE id = $4; `

	result, err := r.db.Exec(query, link.Title, link.URL, link.Desc, id)
	if err != nil {
		return err
	}

	rowsaff, err := result.RowsAffected()
	if rowsaff == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *PostgresRepo) RemoveLink(ctx context.Context, id int) error {
	query := "DELETE FROM links WHERE id = $1"

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsaff, err := result.RowsAffected()
	if rowsaff == 0 {
		return sql.ErrNoRows
	}

	return nil
}
