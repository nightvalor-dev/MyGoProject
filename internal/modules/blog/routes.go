package blog

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, handler *BlogHandler) {
	r.Route("/blogs", func(r chi.Router) {
		r.Get("/", handler.GetAllBlogs)
		r.Post("/", handler.CreateBlog)
		r.Get("/{id}", handler.GetBlogById)
		r.Put("/{id}", handler.UpdateBlog)
		r.Delete("/{id}", handler.DeleteBlog)
	})
}
