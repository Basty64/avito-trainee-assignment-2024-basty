package services

import (
	"avito-trainee-assignment-2024-basty/internal/models"
	"avito-trainee-assignment-2024-basty/internal/repository/postgres"
	"context"
)

type Service struct {
	bannerRepository *postgres.BannersRepository
}

func NewService(repository *postgres.BannersRepository) *Service {
	return &Service{
		bannerRepository: repository,
	}
}

//func GETUserBanner() (models.BannerResponse, error){}
//func GETBanner() ([]models.BannerResponse, error){}

func (s *Service) POSTBanner(ctx context.Context, TagIDS []int, FeatureID int, Content models.Content, IsActive bool) (int, error) {

	return s.bannerRepository.POSTBanner(ctx, TagIDS, FeatureID, Content, IsActive)

}

//func DELETEBanner() error{}
//func PATCHBanner() error{}

//type BannerRepository interface {
//	GETUserBanner() (models.BannerResponse, error)
//	GETBanners() (models.BannerResponse, error)
//	POSTBanner() (int, error)
//	PATCHBanner() error
//	DELETEBanner() error
//}
