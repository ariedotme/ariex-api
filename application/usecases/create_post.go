package usecases

import (
	"time"

	"github.com/ariedotme/ariex-backend/application/services"
	"github.com/ariedotme/ariex-backend/domain/entities"
	"github.com/ariedotme/ariex-backend/shared/utils"
)

type CreatePostUseCase struct {
	PostService *services.PostService
}

func (uc *CreatePostUseCase) Execute(title, content string) (*entities.Post, error) {
	post := &entities.Post{
		ID:              utils.GenerateID(),
		Title:           title,
		NormalizedTitle: utils.NormalizeString(title),
		Content:         content,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	return uc.PostService.CreatePost(post)
}
