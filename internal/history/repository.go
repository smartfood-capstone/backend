package history

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/util"
)

type repository struct {
	db *sqlx.DB
	l  *logrus.Logger
}

type IRepository interface {
	GetAllHistory(ctx context.Context, param getAllParams) ([]History, error)
}

func NewRepository(db *sqlx.DB, l *logrus.Logger) IRepository {
	return &repository{
		db: db,
		l:  l,
	}
}

type getAllParams struct {
	Limit  int
	Offset int
}

func (r *repository) GetAllHistory(ctx context.Context, param getAllParams) ([]History, error) {
	tx, err := r.db.Begin()
	if err != nil {
		r.l.Errorf("error when starting transaction, err: %s", err)
		return nil, err
	}

	query := `select
	dh.id,
	f."name",
	dh."result"->>'data' as category,
	dh.created_at,
	f.description,
	f.image
from
	detection_history dh
inner join foods f on dh."result"->>'data' = f.category WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`

	result := make([]History, 0)
	// For now user_id is hardcoded to 1
	res, err := tx.QueryContext(ctx, query, "1", param.Limit, param.Offset)
	if err != nil {
		r.l.Errorf("error when getting all history, err: %s", err)
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var history History
		err = res.Scan(&history.Id, &history.Name, &history.Category, &history.CreatedAt, &history.Description, &history.Image)
		if err != nil {
			r.l.Errorf("error when scanning history, err: %s", err)
			return nil, err
		}
		result = append(result, history)
	}

	defer util.CommitOrRollback(tx)
	return result, nil
}
