package service

import (
    "main/internal/model"
    "main/internal/store"
)

type BlogService interface {
    Create(model.BlogPost) (model.BlogPost, error)
    GetAll() ([]model.BlogPost, error)
    Update(id int, post model.BlogPost) (model.BlogPost, error)
    Delete(id int) error
}

type blogService struct {
    store store.BlogStore
}

func NewBlogService(s store.BlogStore) BlogService {
    return &blogService{store: s}
}

func (s *blogService) Create(post model.BlogPost) (model.BlogPost, error) {
    return s.store.Create(post)
}

func (s *blogService) GetAll() ([]model.BlogPost, error) {
    return s.store.GetAll()
}

func (s *blogService) Update(id int, post model.BlogPost) (model.BlogPost, error) {
    return s.store.Update(id, post)
}

func (s *blogService) Delete(id int) error {
    return s.store.Delete(id)
}