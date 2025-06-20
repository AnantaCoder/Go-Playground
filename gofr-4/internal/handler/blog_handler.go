package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
    "main/internal/model"
    "main/internal/service"
)

type BlogHandler struct {
    blogService service.BlogService
}

func NewBlogHandler(blogService service.BlogService) *BlogHandler {
    return &BlogHandler{blogService}
}

func (h *BlogHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
    var post model.BlogPost
    _ = json.NewDecoder(r.Body).Decode(&post)
    created, _ := h.blogService.Create(post)
    json.NewEncoder(w).Encode(created)
}

func (h *BlogHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    posts, _ := h.blogService.GetAll()
    json.NewEncoder(w).Encode(posts)
}

func (h *BlogHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/post/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }
    var post model.BlogPost
    _ = json.NewDecoder(r.Body).Decode(&post)
    updated, _ := h.blogService.Update(id, post)
    json.NewEncoder(w).Encode(updated)
}

func (h *BlogHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/post/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }
    _ = h.blogService.Delete(id)
    w.WriteHeader(http.StatusNoContent)
}