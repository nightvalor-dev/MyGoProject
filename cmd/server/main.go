package main

import (
	"Project2-v3/internal/modules/comment"
	"Project2-v3/internal/modules/user"
	"fmt"
	"log"
	"net/http"

	"Project2-v3/api"
	"Project2-v3/config"
	"Project2-v3/internal/modules/blog"
	"Project2-v3/internal/modules/category"
	"Project2-v3/internal/modules/tag"
	db "Project2-v3/internal/shared/db"
)

func main() {
	cfg := config.Load()

	pool, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	blogRepo := blog.NewBlogRepository(pool)
	categoryRepo := category.NewCatRepository(pool)
	tagRepo := tag.NewTagRepository(pool)

	blogService := blog.NewBlogService(blogRepo, categoryRepo, tagRepo)
	categoryService := category.NewCategoryService(categoryRepo)
	tagService := tag.NewTagService(tagRepo)

	blogHandler := blog.NewBlogHandler(blogService)
	categoryHandler := category.NewCategoryHandler(categoryService)
	tagHandler := tag.NewTagHandler(tagService)

	commentRepo := comment.NewCommentRepository(pool)
	commentService := comment.NewCommentService(commentRepo)
	commentHandler := comment.NewCommentHandler(commentService)

	userRepo := user.NewUserRepository(pool)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	router := api.NewRouter(blogHandler, categoryHandler, tagHandler, commentHandler, userHandler)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("server running on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
