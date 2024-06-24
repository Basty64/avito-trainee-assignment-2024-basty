package handlers

import (
	"avito-trainee-assignment-2024-basty/internal/services"
	"encoding/json"
	"math/rand"
	"net/http"
)

type GETBannerHandler struct {
	service *services.CreateBannerService
}

func NewGETBannerHandler(s *services.CreateBannerService) *POSTBannerHandler {
	return &POSTBannerHandler{service: s}
}

type GETBannerRequest struct {
	Content  string `json:"content"`
	Surname  string `json:"surname"`
	Nickname string `json:"nickname"`
}

func (P GETBannerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	message := POSTBannerRequest{}
	err := json.NewDecoder(request.Body).Decode(&message)

	print(Reverse(message.Surname), "  ", Reverse(message.Nickname))

	savedMessage := GETMessagesResponse{message.Nickname, rand.Int()}

	response, err := json.Marshal(savedMessage)
	if err != nil {
		print(err)
	}

	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write(response)
	if err != nil {
		print(err)
	}

}

type GETMessagesResponse struct {
	Content string
	ID      int
}
