package user

import (
	"context"

	"github.com/alexkazantsev/templ-api/domain"
	"github.com/google/uuid"
)

type UserService interface {
	FindOne(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type UserServiceImpl struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (u *UserServiceImpl) FindOne(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return u.repository.FindOne(ctx, id)
}
