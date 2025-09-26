package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nelsonfrank/backend-api-go/internal/services"
	v1 "github.com/nelsonfrank/backend-api-go/internal/transport/http/v1"
	v2 "github.com/nelsonfrank/backend-api-go/internal/transport/http/v2"
)

// API holds all dependencies needed for routing
type API struct {
	UserService *services.UserService
}

func NewAPI(userService *services.UserService) *API {
	return &API{
		UserService: userService,
	}
}

// Router builds and returns the chi router with versioned routes
func (api *API) Router() http.Handler {
	r := chi.NewRouter()

	// ===== V1 =====
	userHandlerV1 := v1.NewUserHandler(api.UserService)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Mount("/", userHandlerV1.Routes())
		})
	})

	// ===== V2 =====
	userHandlerV2 := v2.NewUserHandler(api.UserService)

	r.Route("/api/v2", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Mount("/", userHandlerV2.Routes())
		})
	})

	return r
}
