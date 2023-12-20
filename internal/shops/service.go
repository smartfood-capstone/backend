package shops

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/sirupsen/logrus"
)

type service struct {
	r IRepository
	l *logrus.Logger
}

type IService interface {
	GetAll(ctx context.Context, p getAllRepoParams) ([]Shop, error)
	GetDetail(ctx context.Context, id int) (ShopDetail, error)
	Create(ctx context.Context, f Shop) (Shop, error)
	Update(ctx context.Context, f Shop, id int) (Shop, error)
	Delete(ctx context.Context, id int) (Shop, error)
}

func NewService(r IRepository, l *logrus.Logger) IService {
	return &service{
		r: r,
		l: l,
	}
}

func (s *service) GetAll(ctx context.Context, p getAllRepoParams) ([]Shop, error) {
	result := make([]Shop, 0)
	result, err := s.r.GetAll(ctx, p)
	if err != nil {
		s.l.Errorf("error when getting all Shops, name: %s, limit: %d, offset: %d, err: %s", p.Name, p.Limit, p.Offset, err)
		return result, err
	}

	return result, nil
}

func (s *service) GetDetail(ctx context.Context, id int) (ShopDetail, error) {
	var result ShopDetail
	shop, err := s.r.GetDetail(ctx, id)
	if err != nil {
		s.l.Errorf("error when getting detail id: %d, err: %s", id, err)
		return result, err
	}

	result.GmapsLink = shop.GmapsLink
	result.Id = shop.Id
	result.Image = shop.Image
	result.Location = shop.Location
	result.Name = shop.Name
	result.Latitude = shop.Latitude
	result.Longitude = shop.Longitude

	// convert id to string
	idstring := strconv.Itoa(id)

	foods, err := s.r.GetFoodsByShopId(ctx, idstring)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		s.l.Errorf("error when getting foods by shop id: %d, err: %s", id, err)
		return result, err
	}

	result.Foods = foods

	return result, nil
}

func (s *service) Create(ctx context.Context, sh Shop) (Shop, error) {
	var result Shop
	result, err := s.r.Create(ctx, sh)
	if err != nil {
		s.l.Errorf("error when inserting to database, name: %s, location: %s, gmpas_link: %s, latitude: %f, longitude: %f, image: %s, err: %s", sh.Name, sh.Location, sh.GmapsLink, sh.Latitude, sh.Longitude, sh.Image, err)
		return result, err
	}

	return result, nil
}

func (s *service) Update(ctx context.Context, sh Shop, id int) (Shop, error) {
	var result Shop
	result, err := s.r.Update(ctx, sh, id)
	if err != nil {
		s.l.Errorf("error when updating to database, name: %s, location: %s, gmpas_link: %s, latitude: %f, longitude: %f, image: %s, err: %s", sh.Name, sh.Location, sh.GmapsLink, sh.Latitude, sh.Longitude, sh.Image, err)
		return result, err
	}

	return result, nil
}

func (s *service) Delete(ctx context.Context, id int) (Shop, error) {
	var result Shop
	result, err := s.r.Delete(ctx, id)
	if err != nil {
		s.l.Errorf("error when deleting database id: %d, err: %s", id, err)
		return result, err
	}
	return result, nil
}
