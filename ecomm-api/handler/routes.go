package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	router = chi.NewRouter()

	router.Route("/products", func(router chi.Router) {
		router.Post("/", handler.createProduct)
		router.Get("/", handler.listProducts)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", handler.getroduct)
			router.Patch("/", handler.updateProduct)
			router.Delete("/", handler.deleteProduct)
		})
	})
	return router
}

func Start(address string) error {
	return http.ListenAndServe(address, router)
}
