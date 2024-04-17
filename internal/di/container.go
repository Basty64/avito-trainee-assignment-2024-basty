package di

import (
	"avito-trainee-assignment-2024-basty/internal/handlers"
	"avito-trainee-assignment-2024-basty/internal/models"
	"avito-trainee-assignment-2024-basty/internal/services"
	"github.com/gorilla/mux"
	"net/http"
)

type Container struct {
	err    error
	router http.Handler

	secretKey string

	postBannerHandler *handlers.POSTBannerHandler
	getBannerHandler  *handlers.GETBannerHandler

	createBanner *services.CreateBannerService

	bannerRepository models.BannerRepository
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) HTTPRouter() http.Handler {
	router := mux.NewRouter()

	router.Handle("/banner", c.POSTBannerHandler()).Methods(http.MethodPost)
	//router.Handle("/banner", handler).Methods("GET")
	//router.Handle("/user_banner", handler).Methods("GET")
	//router.Handle("/banner/{id}", handler).Methods("PATCH")
	//router.Handle("/banner/{id}", handler).Methods("DELETE")

	c.router = router
	return c.router
}

func (c *Container) POSTBannerHandler() *handlers.POSTBannerHandler {
	if c.postBannerHandler == nil {
		c.postBannerHandler = handlers.NewPOSTBannerHandler(c.createBanner)
	}
	return c.postBannerHandler
}

func (c *Container) BannersRepository() models.BannerRepository {
	if c.bannerRepository == nil {
		//c.bannerRepository = po -- Сделать репозиторий!!
	}
	return c.bannerRepository
}
