package database

import "github.com/ariedotme/ariex-backend/domain/entities"

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) FindByUsername(username string) (*entities.User, error) {
	var users []entities.User
	err := Client.DB.From("users").Select("*").Eq("username", username).Execute(&users)
	if err != nil || len(users) == 0 {
		return nil, err
	}
	return &users[0], nil
}

func (r *UserRepositoryImpl) Save(user *entities.User) error {
	var results []entities.User
	err := Client.DB.From("users").Insert(user).Execute(&results)
	return err
}
