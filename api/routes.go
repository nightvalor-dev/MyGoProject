package api

import (
	"Project2-v3/internal/modules/blog"
	"Project2-v3/internal/modules/category"
	"Project2-v3/internal/modules/comment"
	"Project2-v3/internal/modules/tag"
	"Project2-v3/internal/modules/user"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(blogHandler *blog.BlogHandler,
	categoryHandler *category.CategoryHandler,
	tagHandler *tag.TagHandler,
	commentHandler *comment.CommentHandler,
	userHandler *user.UserHandler) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	blog.RegisterRoutes(r, blogHandler)
	category.RegisterRoutes(r, categoryHandler)
	tag.RegisterRoutes(r, tagHandler)
	comment.RegisterRoutes(r, commentHandler)
	user.RegisterRoutes(r, userHandler)

	return r
}
