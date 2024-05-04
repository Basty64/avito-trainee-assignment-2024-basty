package postgres

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BannersRepository struct {
	pool *pgxpool.Pool
}

func NewBannersRepository(pool *pgxpool.Pool) *BannersRepository {
	return &BannersRepository{pool: pool}
}

var errNotImplemented = errors.New("I dont know postgres")

//func (b *BannersRepository) GETUserBanner() (models.BannerResponse, error) { return nil, nil }
//func (b *BannersRepository) GETBanners() ([]models.BannerResponse, error) { return nil, nil }
//func (b *BannersRepository) POSTBanner() (int, error){return nil, nil}
//func (b *BannersRepository) PATCHBanner() error{}
//func (b *BannersRepository) DELETEBanner() error{}
