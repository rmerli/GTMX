// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: category.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (name, section_id)
VALUES ($1, $2)
RETURNING id, name, section_id
`

type CreateCategoryParams struct {
	Name      string
	SectionID pgtype.Int8
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory, arg.Name, arg.SectionID)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.SectionID)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, section_id FROM categories
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRow(ctx, getCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.SectionID)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, section_id FROM categories
ORDER BY name
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.Query(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.SectionID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories
  set name = $2, section_id = $3
  WHERE id = $1
RETURNING id, name, section_id
`

type UpdateCategoryParams struct {
	ID        int64
	Name      string
	SectionID pgtype.Int8
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, updateCategory, arg.ID, arg.Name, arg.SectionID)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.SectionID)
	return i, err
}
