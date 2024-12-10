package services

import (
	"errors"

	"github.com/ariedotme/ariex-backend/domain/entities"
	"github.com/ariedotme/ariex-backend/domain/repositories"
	"github.com/ariedotme/ariex-backend/shared/utils"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func (s *UserService) CreateUser(username, password string) (*entities.User, error) {

	existingUser, _ := s.UserRepository.FindByUsername(username)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:       utils.GenerateID(),
		Username: username,
		Password: hashedPassword,
	}

	er := s.UserRepository.Save(user)

	if er != nil {
		return nil, er
	}

	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (*entities.User, error) {
	return s.UserRepository.FindByUsername(username)
}
