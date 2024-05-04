package services

import (
	"avito-trainee-assignment-2024-basty/internal/models"
	"avito-trainee-assignment-2024-basty/internal/repository/postgres"
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

func (s *Service) POSTBanner(Tag_IDS []int, feature_id int, content models.Content, isActive bool) (int, error) {

	return 0, nil
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
