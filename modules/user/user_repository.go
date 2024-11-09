package user

import (
	"context"

	"github.com/alexkazantsev/templ-api/domain"
	"github.com/alexkazantsev/templ-api/modules/database/storage"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindOne(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type UserRepositoryImpl struct {
	q *storage.Queries
}

func NewUserRepository(q *storage.Queries) UserRepository {
	return &UserRepositoryImpl{
		q: q,
	}
}

func (u *UserRepositoryImpl) FindOne(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.q.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}, nil
}
