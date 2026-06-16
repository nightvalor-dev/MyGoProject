package comment

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, handler *CommentHandler) {
	r.Route("/comments", func(r chi.Router) {
		r.Get("/", handler.GetAllComments)
		r.Post("/", handler.CreateComment)
		r.Get("/{id}", handler.GetCommentById)
		r.Put("/{id}", handler.UpdateComment)
		r.Delete("/{id}", handler.DeleteComment)
	})
}
