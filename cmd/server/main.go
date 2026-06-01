package main

import (
	"fmt"
	"log"
	"net/http"

	"Project2-v3/api"
	"Project2-v3/config"
	"Project2-v3/internal/modules/blog"
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
	blogService := blog.NewBlogService(blogRepo)
	blogHandler := blog.NewBlogHandler(blogService)

	router := api.NewRouter(blogHandler)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("server running on %s", addr)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
