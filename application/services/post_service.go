package services

import (
	"github.com/ariedotme/ariex-backend/domain/entities"
	"github.com/ariedotme/ariex-backend/domain/repositories"
)

type PostService struct {
	PostRepository repositories.PostRepository
}

func (s *PostService) CreatePost(post *entities.Post) (*entities.Post, error) {
	if err := s.PostRepository.Save(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) GetPostByNormalizedTitle(normalizedTitle string) (*entities.Post, error) {
	post, err := s.PostRepository.FindByNormalizedTitle(normalizedTitle)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) GetPostByID(id string) (*entities.Post, error) {
	post, err := s.PostRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) GetPosts() ([]*entities.Post, error) {
	posts, err := s.PostRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) UpdatePost(post *entities.Post) error {
	err := s.PostRepository.Update(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeletePost(id string) error {
	err := s.PostRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
