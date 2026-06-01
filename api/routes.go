package api

import (
	"Project2-v3/internal/modules/blog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(blogHandler *blog.BlogHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	blog.RegisterRoutes(r, blogHandler)

	return r
}
