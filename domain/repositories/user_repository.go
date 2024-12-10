package repositories

import "github.com/ariedotme/ariex-backend/domain/entities"

type UserRepository interface {
	FindByUsername(username string) (*entities.User, error)
	Save(user *entities.User) error
}
