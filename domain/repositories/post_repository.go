package repositories

import "github.com/ariedotme/ariex-backend/domain/entities"

type PostRepository interface {
	FindAll() ([]*entities.Post, error)
	FindByID(id string) (*entities.Post, error)
	FindByNormalizedTitle(normalizedTitle string) (*entities.Post, error)
	Save(post *entities.Post) error
	Update(post *entities.Post) error
	Delete(id string) error
}
