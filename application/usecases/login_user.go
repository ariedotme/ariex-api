package usecases

import (
	"errors"

	"github.com/ariedotme/ariex-backend/application/services"
	"github.com/ariedotme/ariex-backend/shared/utils"
)

type LoginUserUseCase struct {
	UserService *services.UserService
}

func (uc *LoginUserUseCase) Execute(username, password string) (string, error) {
	user, err := uc.UserService.GetUserByUsername(username)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
