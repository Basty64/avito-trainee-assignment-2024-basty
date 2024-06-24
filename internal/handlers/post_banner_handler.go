package handlers

import (
	"avito-trainee-assignment-2024-basty/internal/models"
	"avito-trainee-assignment-2024-basty/internal/services"
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	POSTBanner(ctx context.Context, Tag_IDS []int, feature_id int, content models.Content, isActive bool) (int, error)
}

type Handlers struct {
	service *services.Service
}

func NewHandlers(s *services.Service) *Handlers {
	return &Handlers{service: s}
}

func POSTBannerHandler(service Service) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		POSTBannerRequest := models.POSTBannerRequest{}

		err := json.NewDecoder(request.Body).Decode(&POSTBannerRequest)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		response, err := service.POSTBanner(
			context.Background(),
			POSTBannerRequest.TagIDS,
			POSTBannerRequest.FeatureID,
			POSTBannerRequest.Content,
			POSTBannerRequest.IsActive)

		if err != nil {
			print(err)
		}

		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			print(err)
		}

		writer.WriteHeader(http.StatusCreated)

		//writer.WriteHeader(http.StatusCreated)
		//_, err = writer.Write(response)
		//if err != nil {
		//	print(err)
		//}
	}
}
