package route

import (
	"net/http"

	handlers "for_learning/src/interface/rest/handlers/pickup"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func PickupRouter(h handlers.BooksHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.Create)

	return r
}
