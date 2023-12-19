package predict

import (
	"context"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/foods"
	"github.com/smartfood-capstone/backend/internal/util"
)

type repository struct {
	db *sqlx.DB
	l  *logrus.Logger
}

type IRepository interface {
	InsertHistory(ctx context.Context, history History) error
	GetFoodDetailByCategory(ctx context.Context, category string) (foods.Food, error)
}

func NewRepository(db *sqlx.DB, l *logrus.Logger) IRepository {
	return &repository{
		db: db,
		l:  l,
	}
}

func (r *repository) InsertHistory(ctx context.Context, history History) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.l.Errorf("error when starting transaction, err: %s", err)
		return err
	}

	query := `INSERT INTO detection_history (user_id, result) VALUES ($1, $2)`

	marshaledResult, err := json.Marshal(history.Result)
	if err != nil {
		r.l.Errorf("error when marshaling result, err: %s", err)
		return err
	}

	_, err = tx.ExecContext(ctx, query, history.UserId, marshaledResult)

	if err != nil {
		r.l.Errorf("error when inserting history, err: %s", err)
		return err
	}

	defer util.CommitOrRollback(tx)
	return nil
}

func (r *repository) GetFoodDetailByCategory(ctx context.Context, category string) (foods.Food, error) {
	var result foods.Food
	query := `
    SELECT * FROM foods f
    WHERE category = $1
  `

	err := r.db.GetContext(ctx, &result, query, category)
	if err != nil {
		r.l.Errorf("error when getting detail id: %s, err: %s", category, err)
		return result, err
	}

	return result, nil
}
