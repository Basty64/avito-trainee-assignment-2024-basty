package services

import "avito-trainee-assignment-2024-basty/internal/models"

type CreateBannerService struct {
	bannerRepository models.BannerRepository
}

func NewCreateBannerService(repository models.BannerRepository) *CreateBannerService {
	return &CreateBannerService{
		bannerRepository: repository,
	}
}
