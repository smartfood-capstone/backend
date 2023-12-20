package shops

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
	GetAll(ctx context.Context, p getAllRepoParams) ([]Shop, error)
	GetDetail(ctx context.Context, id int) (Shop, error)
	Create(ctx context.Context, f Shop) (Shop, error)
	Update(ctx context.Context, f Shop, id int) (Shop, error)
	Delete(ctx context.Context, id int) (Shop, error)
	GetFoodsByShopId(ctx context.Context, id string) ([]FoodShop, error)
}

func NewRepository(db *sqlx.DB, l *logrus.Logger) IRepository {
	return &repository{
		db: db,
		l:  l,
	}
}

type getAllRepoParams struct {
	Name   string
	Limit  int
	Offset int
}

func (r *repository) GetAll(ctx context.Context, p getAllRepoParams) ([]Shop, error) {
	result := make([]Shop, 0)
	query := `
  SELECT * FROM shops s
  WHERE "name" ilike '%' || $1 || '%'
  LIMIT $2 OFFSET $3
  `

	err := r.db.SelectContext(ctx, &result, query, p.Name, p.Limit, p.Offset)
	if err != nil {
		r.l.Errorf("error when getting all shops, name: %s, limit: %d, offset: %d, err: %s", p.Name, p.Limit, p.Offset, err)
		return result, err
	}

	return result, nil
}

func (r *repository) GetDetail(ctx context.Context, id int) (Shop, error) {
	var result Shop
	query := `
    SELECT * FROM shops s
    WHERE id = $1
  `

	err := r.db.GetContext(ctx, &result, query, id)
	if err != nil {
		r.l.Errorf("error when getting detail id: %d, err: %s", id, err)
		return result, err
	}

	return result, nil
}

func (r *repository) Create(ctx context.Context, s Shop) (Shop, error) {
	var result Shop
	query := `
    INSERT INTO "shops" (name, location, gmaps_link, latitude, longitude, image)
    VALUES ($1,$2,$3,$4,$5,$6)
    RETURNING *
  `

	err := r.db.GetContext(ctx, &result, query, s.Name, s.Location, s.GmapsLink, s.Latitude, s.Longitude, s.Image)
	if err != nil {
		r.l.Errorf("error when inserting to database, name: %s, location: %s, gmpas_link: %s, latitude: %f, longitude: %f, image: %s, err: %s", s.Name, s.Location, s.GmapsLink, s.Latitude, s.Longitude, s.Image, err)
		return result, err
	}

	return result, nil
}

func (r *repository) Update(ctx context.Context, s Shop, id int) (Shop, error) {
	var result Shop
	existingShop, err := r.GetDetail(ctx, id)
	if err != nil {
		r.l.Errorf("error when getting id from database, id: %d, err: %s", id, err)
		return result, err
	}

	if s.Name != "" {
		existingShop.Name = s.Name
	}

	if s.Location != "" {
		existingShop.Location = s.Location
	}

	if s.GmapsLink != "" {
		existingShop.GmapsLink = s.GmapsLink
	}

	if s.Image != "" {
		existingShop.Image = s.Image
	}

	if s.Latitude != 0 {
		existingShop.Latitude = s.Latitude
	}

	if s.Longitude != 0 {
		existingShop.Longitude = s.Longitude
	}

	query := `
    UPDATE shops
    SET "name" = $1, location = $2, gmaps_link = $3, latitude = $4, longitude = $5, image = $6
    WHERE id = $7
    RETURNING *
  `

	err = r.db.GetContext(ctx, &result, query, existingShop.Name, existingShop.Location, existingShop.GmapsLink, existingShop.Latitude, existingShop.Longitude, existingShop.Image, id)

	if err != nil {
		r.l.Errorf("error when updating to database, name: %s, location: %s, gmpas_link: %s, latitude: %f, longitude: %f, image: %s, err: %s", s.Name, s.Location, s.GmapsLink, s.Latitude, s.Longitude, s.Image, err)
		return result, err
	}

	return result, nil
}

func (r *repository) Delete(ctx context.Context, id int) (Shop, error) {
	var result Shop
	_, err := r.GetDetail(ctx, id)
	if err != nil {
		r.l.Errorf("error when getting id from database, id: %d, err: %s", id, err)
		return result, err
	}

	query := `
    DELETE FROM shops
    WHERE id = $1
    RETURNING *
  `

	err = r.db.GetContext(ctx, &result, query, id)
	if err != nil {
		r.l.Errorf("error when deleting database id: %d, err: %s", id, err)
		return result, err
	}

	return result, nil
}

func (r *repository) GetFoodsByShopId(ctx context.Context, id string) ([]FoodShop, error) {
	result := make([]FoodShop, 0)
	query := `
	SELECT f.id, f.name, fs.price FROM foods f
	INNER JOIN shop_foods fs ON f.id = fs.food_id
	WHERE fs.shop_id = $1
	`

	err := r.db.SelectContext(ctx, &result, query, id)
	if err != nil {
		r.l.Errorf("error when getting all foods by shop id, id: %d, err: %s", id, err)
		return result, err
	}

	return result, nil
}
