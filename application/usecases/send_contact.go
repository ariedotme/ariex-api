package usecases

import (
	"github.com/ariedotme/ariex-backend/application/services"
	"github.com/ariedotme/ariex-backend/domain/entities"
)

type SendContactUseCase struct {
	EmailService *services.EmailService
}

func (uc *SendContactUseCase) Execute(contact *entities.Contact) error {
	return uc.EmailService.Send(contact.Email, "New Contact Message", contact.Message)
}
