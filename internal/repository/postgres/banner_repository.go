package postgres

import (
	"avito-trainee-assignment-2024-basty/internal/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
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

func (b *BannersRepository) POSTBanner(ctx context.Context, TagIDS []int, FeatureID int, Content models.Content, IsActive bool) (int, error) {

	var banners_id int

	POSTquery := `INSERT INTO banners.public.banners (tag_ids, feature_id, content, is_active)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING banner_id
	`
	err := b.pool.QueryRow(ctx, POSTquery, TagIDS, FeatureID, Content, IsActive).Scan(&banners_id)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return banners_id, nil

}

//func (b *BannersRepository) PATCHBanner() error{}
//func (b *BannersRepository) DELETEBanner() error{}
