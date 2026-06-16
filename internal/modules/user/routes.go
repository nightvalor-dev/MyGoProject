package user

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, handler *UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetAllUsers)
		r.Post("/", handler.CreateUser)
		r.Get("/{id}", handler.GetUserById)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})
}
