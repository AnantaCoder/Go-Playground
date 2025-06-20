package store

import (
	"context"
	"encoding/json"
	"fmt"
	"main/internal/model"

	"github.com/redis/go-redis/v9"
)

type BlogStore interface {
	Create(model.BlogPost)(model.BlogPost,error)
	GetAll()([]model.BlogPost,error)
	Update(id int ,post model.BlogPost)(model.BlogPost,error)
	Delete(id int)error
}


type redisBlogStore struct{
	client  *redis.Client
	ctx  context.Context
}

func NewRedisBlogStore(addr, password string, db int) BlogStore {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })
    return &redisBlogStore{client: rdb, ctx: context.Background()}
}
func (s *redisBlogStore) key(id int) string {
    return fmt.Sprintf("blog:%d", id)
}

func (s *redisBlogStore) Create(post model.BlogPost) (model.BlogPost, error) {
    id, err := s.client.Incr(s.ctx, "blog:next-id").Result()
    if err != nil {
        return post, err
    }
    post.ID = int(id)
    data, _ := json.Marshal(post)
    err = s.client.Set(s.ctx, s.key(post.ID), data, 0).Err()
    return post, err
}

func (s *redisBlogStore) GetAll() ([]model.BlogPost, error) {
    keys, err := s.client.Keys(s.ctx, "blog:*").Result()
    if err != nil {
        return nil, err
    }
    var posts []model.BlogPost
    for _, key := range keys {
        if key == "blog:next-id" {
            continue
        }
        val, err := s.client.Get(s.ctx, key).Result()
        if err == nil {
            var post model.BlogPost
            json.Unmarshal([]byte(val), &post)
            posts = append(posts, post)
        }
    }
    return posts, nil
}

func (s *redisBlogStore) Update(id int, post model.BlogPost) (model.BlogPost, error) {
    post.ID = id
    data, _ := json.Marshal(post)
    err := s.client.Set(s.ctx, s.key(id), data, 0).Err()
    return post, err
}

func (s *redisBlogStore) Delete(id int) error {
    return s.client.Del(s.ctx, s.key(id)).Err()
}