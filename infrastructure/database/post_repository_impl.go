package database

import (
	"fmt"

	"github.com/ariedotme/ariex-backend/domain/entities"
)

type PostRepositoryImpl struct{}

func (r *PostRepositoryImpl) FindAll() ([]*entities.Post, error) {
	var posts []*entities.Post
	err := Client.DB.From("posts").Select("*").Execute(&posts)
	if err != nil {
		return nil, err
	}

	if posts == nil {
		posts = []*entities.Post{}
	}

	return posts, err
}

func (r *PostRepositoryImpl) FindByID(id string) (*entities.Post, error) {
	var posts []entities.Post
	err := Client.DB.From("posts").Select("*").Eq("id", id).Execute(&posts)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, fmt.Errorf("no post found with  id: %s", id)
	}
	return &posts[0], nil
}

func (r *PostRepositoryImpl) FindByNormalizedTitle(normalizedTitle string) (*entities.Post, error) {
	var posts []entities.Post
	err := Client.DB.From("posts").Select("*").Eq("normalized_title", normalizedTitle).Execute(&posts)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, fmt.Errorf("no post found with  normalized title: %s", normalizedTitle)
	}
	return &posts[0], nil
}

func (r *PostRepositoryImpl) Save(post *entities.Post) error {
	var results []entities.Post
	err := Client.DB.From("posts").Insert(post).Execute(&results)
	return err
}

func (r *PostRepositoryImpl) Update(post *entities.Post) error {
	var results []entities.Post
	err := Client.DB.From("posts").Update(post).Eq("id", post.ID).Execute(&results)
	return err
}

func (r *PostRepositoryImpl) Delete(id string) error {
	var results []entities.Post
	err := Client.DB.From("posts").Delete().Eq("id", id).Execute(&results)
	return err
}
