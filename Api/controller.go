package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"

)

func (a *App) Router() http.Handler {
	fmt.Println("Reached here")
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/products", a.handleGetProducts)
		r.Post("/product", a.handleCreateProduct)
		r.Post("/product/{id}", a.handleGetProduct)
		r.Put("/product/{id}", a.handleUpdateProduct)
		r.Delete("/product/{id}", a.handleDeleteProduct)
		r.Get("/store/{id}", a.handleStoreProduct)
		r.Post("/store/{id}", a.handleInsertProduct)

	})
	return r
}


