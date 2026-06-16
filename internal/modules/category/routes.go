package category

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, handler *CategoryHandler) {
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", handler.GetAllCategory)
		r.Post("/", handler.CreateCategory)
		r.Get("/{id}", handler.GetCategoryById)
		r.Put("/{id}", handler.UpdateCategory)
		r.Delete("/{id}", handler.DeleteCategory)
	})
}
