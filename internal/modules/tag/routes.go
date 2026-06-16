package tag

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, handler *TagHandler) {
	r.Route("/tags", func(r chi.Router) {
		r.Get("/", handler.GetAllTags)
		r.Post("/", handler.CreateTag)
		r.Get("/{id}", handler.GetTagById)
		r.Put("/{id}", handler.UpdateTag)
		r.Delete("/{id}", handler.DeleteTag)
	})
}
