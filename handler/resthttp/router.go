package resthttp

import (
	"github.com/go-chi/chi"
)

type RouterDependencies struct {
	PS ProductService
}

func NewRoutes(rd RouterDependencies) *chi.Mux {
	router := chi.NewRouter()

	ph := newProductHandler(rd.PS)
	router.Post("/update-all", ph.UpdateAllProduct)

	return router
}
