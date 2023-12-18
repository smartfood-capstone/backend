package foods

import (
	"context"

	"github.com/sirupsen/logrus"
)

type service struct {
	r IRepository
	l *logrus.Logger
}

type IService interface {
	GetAll(ctx context.Context, p getAllRepoParams) ([]Food, error)
	GetDetail(ctx context.Context, id int) (Food, error)
	Create(ctx context.Context, f Food) (Food, error)
	Update(ctx context.Context, f Food, id int) (Food, error)
	Delete(ctx context.Context, id int) (Food, error)
}

func NewService(r IRepository, l *logrus.Logger) IService {
	return &service{
		r: r,
		l: l,
	}
}

func (s *service) GetAll(ctx context.Context, p getAllRepoParams) ([]Food, error) {
	var result = make([]Food, 0)
	result, err := s.r.GetAll(ctx, p)
	if err != nil {
		s.l.Errorf("error when getting all foods, name: %s, limit: %d, offset: %d, err: %s", p.Name, p.Limit, p.Offset, err)
		return result, err
	}

	return result, nil
}

func (s *service) GetDetail(ctx context.Context, id int) (Food, error) {
	var result Food
	result, err := s.r.GetDetail(ctx, id)
	if err != nil {
		s.l.Errorf("error when getting detail id: %d, err: %s", id, err)
		return result, err
	}

	return result, nil
}

func (s *service) Create(ctx context.Context, f Food) (Food, error) {
	var result Food
	result, err := s.r.Create(ctx, f)
	if err != nil {
		s.l.Errorf("error when creating database, name: %s, desc: %s, category: %s, image: %s, err: %s", f.Name, f.Description, f.Category, f.Image, err)
		return result, err
	}

	return result, nil
}

func (s *service) Update(ctx context.Context, f Food, id int) (Food, error) {
	var result Food
	result, err := s.r.Update(ctx, f, id)
	if err != nil {
		s.l.Errorf("error when updating database, name: %s, desc: %s, category: %s, image: %s, err: %s", f.Name, f.Description, f.Category, f.Image, err)
		return result, err
	}

	return result, nil
}

func (s *service) Delete(ctx context.Context, id int) (Food, error) {
	var result Food
	result, err := s.r.Delete(ctx, id)
	if err != nil {
		s.l.Errorf("error when deleting database id: %d, err: %s", id, err)
		return result, err
	}
	return result, nil
}
