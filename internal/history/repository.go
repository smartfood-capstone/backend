package history

// import (
// 	"context"
// 	"encoding/json"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/sirupsen/logrus"
// 	"github.com/smartfood-capstone/backend/internal/util"
// )

// type repository struct {
// 	db *sqlx.DB
// 	l  *logrus.Logger
// }

// type IRepository interface {
// 	GetAllHistory(ctx context.Context) error
// }

// func NewRepository(db *sqlx.DB, l *logrus.Logger) IRepository {
// 	return &repository{
// 		db: db,
// 		l:  l,
// 	}
// }

// func (r *repository) GetAllHistory(ctx context.Context) ([]History, error) {
// 	tx, err := r.db.Begin()
// 	if err != nil {
// 		r.l.Errorf("error when starting transaction, err: %s", err)
// 		return nil, err
// 	}

// 	query := `SELECT  FROM detection_history`

// 	defer util.CommitOrRollback(tx)
// 	return nil
// }
