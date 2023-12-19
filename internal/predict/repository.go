package predict

import (
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
	return &repository{
		db: db,
		l:  l,
	}
}

func (r *repository) InsertHistory() error {

	

	return nil
}
