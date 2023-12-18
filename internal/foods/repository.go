package foods

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type repository struct {
	db *sqlx.DB
	l  *logrus.Logger
}

type IRepository interface {
}

func NewRepository(db *sqlx.DB, l *logrus.Logger) IRepository {
	return repository{
		db: db,
		l:  l,
	}
}

type GetAllParams struct {
	Name   string
	Limit  int
	Offset int
}

func (r *repository) GetAll(ctx context.Context, p GetAllParams) ([]Food, error) {
	var result []Food
	query := `
  SELECT * FROM foods f
  WHERE "name" ilike '%' || $1 || '%'
  LIMIT $2 OFFSET $3
  `

	err := r.db.SelectContext(ctx, &result, query, p.Name, p.Limit, p.Offset)
	if err != nil {
		r.l.Errorf("error when getting all foods, name: %s, limit: %d, offset: %d, err: %s", p.Name, p.Limit, p.Offset, err)
		return result, err
	}

	return result, nil
}

func (r *repository) GetDetail(ctx context.Context, id int) (Food, error) {
	var result Food
	query := `
    SELECT * FROM foods f
    WHERE id = $1
  `

	err := r.db.GetContext(ctx, &result, query, id)
	if err != nil {
		r.l.Errorf("error when getting detail id: %d, err: %s", id, err)
		return result, err
	}

	return result, nil
}

func (r *repository) Create(ctx context.Context, f Food) (Food, error) {
	var result Food
	query := `
    INSERT INTO "foods" (name, description, category, image)
    VALUES ($1,$2,$3,$4) RETURNING *
  `

	err := r.db.GetContext(ctx, &result, query)
	if err != nil {
		r.l.Errorf("error when creating database, name: %s, desc: %s, category: %s, image: %s, err: %s", f.Name, f.Description, f.Category, f.Image, err)
		return result, err
	}

	return result, nil
}
func (r *repository) Update() {}
func (r *repository) Delete() {}
