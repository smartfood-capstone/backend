package history

import (
	"context"

	"github.com/sirupsen/logrus"
)

type service struct {
	r IRepository
	l *logrus.Logger
}

type IService interface {
	GetAll(ctx context.Context, p getAllParams) ([]History, error)
}

func NewService(r IRepository, l *logrus.Logger) IService {
	return &service{
		r: r,
		l: l,
	}
}

func (s *service) GetAll(ctx context.Context, p getAllParams) ([]History, error) {
	var result = make([]History, 0)
	result, err := s.r.GetAllHistory(ctx, p)
	if err != nil {
		s.l.Errorf("error when getting all history, limit: %d, offset: %d, err: %s", p.Limit, p.Offset, err)
		return result, err
	}

	return result, nil
}
