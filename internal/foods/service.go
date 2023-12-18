package foods

import (
	"context"

	"github.com/sirupsen/logrus"
)

type service struct {
	r IRepository
	l *logrus.Logger
}

type IService interface{}

func NewService(r IRepository, l *logrus.Logger) IService {
	return &service{
		r: r,
		l: l,
	}
}

type getAllServiceParams struct {
	Name  string
	Limit int
	Page  int
}

func (s *service) GetAll(ctx context.Context, p getAllServiceParams) ([]Food, error) {
	var result = make([]Food, 0)

	offset := (p.Page - 1) * 20

	params := getAllRepoParams{
		Name:   p.Name,
		Limit:  p.Limit,
		Offset: offset,
	}

	result, err := s.r.GetAll(ctx, params)
	if err != nil {
		s.l.Errorf("error when getting all foods, name: %s, limit: %d, page: %d, err: %s", p.Name, p.Limit, p.Page, err)
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
