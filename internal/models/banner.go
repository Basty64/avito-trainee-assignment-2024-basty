package models

import "github.com/gofrs/uuid"

type Banner struct {
	banner_id  uuid.UUID `json:"banner_Id,omitempty"`
	tag_id     []int
	feature_id int
	content    string
	is_active  bool
}

func NewBanner(tag_id []int, feature_id int, content string) *Banner {
	return &Banner{
		banner_id:  uuid.Must(uuid.NewV7()),
		tag_id:     tag_id,
		feature_id: feature_id,
		content:    content,
		is_active:  false,
	}
}

func (b *Banner) ISActive() {
	b.is_active = true
}

type BannerRepository interface {
	Save(banner *Banner) error
}
