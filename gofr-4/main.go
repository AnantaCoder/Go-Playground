package main

import (
	"fmt"
	"log"
	"main/internal/handler"
	"main/internal/service"
	"main/internal/store"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {


	err := godotenv.Load(".env")
    	if err != nil {
    		log.Fatalf("Error loading .env file: %s", err)
    	}
    redisAddr := os.Getenv("REDIS_ADDR")
    redisPassword := os.Getenv("REDIS_PASSWORD")
    redisDB := 0

    redisStore := store.NewRedisBlogStore(redisAddr,redisPassword, redisDB)
    service := service.NewBlogService(redisStore)
    blogHandler := handler.NewBlogHandler(service)

    http.HandleFunc("/post", blogHandler.CreatePost)
    http.HandleFunc("/posts", blogHandler.GetAll)
    http.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPut:
            blogHandler.UpdatePost(w, r)
        case http.MethodDelete:
            blogHandler.DeletePost(w, r)
        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    })

    fmt.Println("serving on 8000")
    http.ListenAndServe(":8000", nil)
}