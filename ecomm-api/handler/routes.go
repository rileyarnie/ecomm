package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var router *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	router = chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/products", func(router chi.Router) {
		router.Post("/", handler.createProduct)
		router.Get("/", handler.listProducts)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", handler.getroduct)
			router.Patch("/", handler.updateProduct)
			router.Delete("/", handler.deleteProduct)
		})
	})
	router.Route("/orders", func(router chi.Router) {
		router.Post("/", handler.createOrder)
		router.Get("/", handler.listOrders)

		router.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getOrder)
			r.Delete("/", handler.deleteOrder)
		})
	})
	router.Route("/users", func(router chi.Router) {
		router.Post("/", handler.createUser)
		router.Get("/", handler.listUsers)
		router.Patch("/", handler.updateUser)

		router.Route("/{id}", func(router chi.Router) {
			router.Delete("/", handler.deleteUser)
		})

		router.Route("/login", func(router chi.Router) {
			router.Post("/", handler.loginUser)
		})

		router.Route("/logout", func(router chi.Router) {
			router.Post("/", handler.logoutUser)
		})
	})

	router.Route("/tokens", func(router chi.Router) {
		router.Route("/renew", func(router chi.Router) {
			router.Post("/", handler.renewAccessToken)
		})

		router.Route("/revoke/{id}", func(router chi.Router) {
			router.Post("/", handler.revokeSession)
		})
	})

	return router
}

func Start(address string) error {
	return http.ListenAndServe(address, router)
}
