package user

import (
	"context"
	"fmt"

	"github.com/alexkazantsev/go-templ-api/domain"
	"github.com/alexkazantsev/go-templ-api/modules/database/storage"
	"github.com/alexkazantsev/go-templ-api/modules/user/dto"
	"github.com/alexkazantsev/go-templ-api/pkg/xerror"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UserRepository interface {
	Exist(context.Context, uuid.UUID) (bool, error)
	FindOne(context.Context, uuid.UUID) (*domain.User, error)
	Create(context.Context, *dto.CreateUserRequest) (*domain.User, error)
	UpdateOne(context.Context, *dto.UpdateUserRequest) (*domain.User, error)
	DeleteOne(context.Context, uuid.UUID) error
}

type UserRepositoryImpl struct {
	q *storage.Queries
}

func NewUserRepository(q *storage.Queries) UserRepository {
	return &UserRepositoryImpl{
		q: q,
	}
}

func (u *UserRepositoryImpl) UpdateOne(ctx context.Context, request *dto.UpdateUserRequest) (*domain.User, error) {
	r, err := u.q.UpdateOne(ctx, storage.UpdateOneParams{
		ID:    request.ID,
		Name:  request.Name,
		Email: request.Email,
	})

	if err != nil {
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

func (u *UserRepositoryImpl) Create(ctx context.Context, request *dto.CreateUserRequest) (*domain.User, error) {
	r, err := u.q.Create(ctx, storage.CreateParams{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return nil, fmt.Errorf("user already exists: %w", xerror.ErrAlreadyExists)
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

func (u *UserRepositoryImpl) Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	return u.q.Exist(ctx, id)
}

func (u *UserRepositoryImpl) DeleteOne(ctx context.Context, id uuid.UUID) error {
	return u.q.Delete(ctx, id)
}
