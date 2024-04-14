package handlers

import (
	"avito-trainee-assignment-2024-basty/internal/services"
	"encoding/json"
	"math/rand"
	"net/http"
)

type POSTBannerHandler struct {
	service *services.CreateBannerService
}

func NewPOSTBannerHandler(s *services.CreateBannerService) *POSTBannerHandler {
	return &POSTBannerHandler{service: s}
}

type POSTBannerRequest struct {
	Content  string `json:"content"`
	Surname  string `json:"surname"`
	Nickname string `json:"nickname"`
}

func (P POSTBannerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	message := POSTBannerRequest{}
	err := json.NewDecoder(request.Body).Decode(&message)

	print(Reverse(message.Surname), "  ", Reverse(message.Nickname))

	savedMessage := POSTMessagesResponse{message.Nickname, rand.Int()}

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

type POSTMessagesResponse struct {
	Content string
	ID      int
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
