package user

import (
	"context"
	"errors"

	"github.com/alexkazantsev/go-templ-api/domain"
	"github.com/alexkazantsev/go-templ-api/modules/database/storage"
	"github.com/alexkazantsev/go-templ-api/modules/user/dto"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

var ErrAlreadyExist = errors.New("user already exists")

type UserRepository interface {
	FindOne(ctx context.Context, id uuid.UUID) (*domain.User, error)
	Create(ctx context.Context, request *dto.CreateUserRequest) (*domain.User, error)
}

type UserRepositoryImpl struct {
	q *storage.Queries
}

func NewUserRepository(q *storage.Queries) UserRepository {
	return &UserRepositoryImpl{
		q: q,
	}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, request *dto.CreateUserRequest) (*domain.User, error) {
	r, err := u.q.Create(ctx, storage.CreateParams{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return nil, ErrAlreadyExist
			}

			return nil, err
		}

		return nil, err
	}

	return &domain.User{
		ID:        r.ID,
		Name:      r.Name,
		Email:     r.Email,
		Password:  r.Password,
		CreatedAt: r.CreatedAt,
	}, nil
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
