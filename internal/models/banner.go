package models

import (
	"time"
)

type POSTBannerRequest struct {
	TagIDS    []int   `json:"tag_ids"`
	FeatureID int     `json:"feature_id"`
	Content   Content `json:"content"`
	IsActive  bool    `json:"is_active"`
}

type BannerResponse struct {
	banner_id  int       `json:"banner_Id,omitempty"`
	tag_ids    []int     `json:"tag_id"`
	feature_id int       `json:"feature_id"`
	content    Content   `json:"content"`
	is_active  bool      `json:"is_active"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}

type Content struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}
