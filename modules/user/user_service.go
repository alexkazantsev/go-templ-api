package user

import (
	"context"
	"fmt"

	"github.com/alexkazantsev/go-templ-api/domain"
	"github.com/alexkazantsev/go-templ-api/modules/core"
	"github.com/alexkazantsev/go-templ-api/modules/user/dto"
	"github.com/alexkazantsev/go-templ-api/pkg/xerror"
	"github.com/google/uuid"
)

type UserService interface {
	FindOne(context.Context, uuid.UUID) (*domain.User, error)
	Create(context.Context, *dto.CreateUserRequest) (*domain.User, error)
	UpdateOne(context.Context, *dto.UpdateUserRequest) (*domain.User, error)
	DeleteOne(context.Context, uuid.UUID) error
}

type UserServiceImpl struct {
	repository  UserRepository
	passwordSrv core.PasswordService
}

func (u *UserServiceImpl) UpdateOne(ctx context.Context, request *dto.UpdateUserRequest) (*domain.User, error) {
	if exist, err := u.repository.Exist(ctx, request.ID); err != nil || !exist {
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("user [%s] does not exists: %w", request.ID, xerror.ErrNotFound)
	}

	return u.repository.UpdateOne(ctx, request)
}

func (u *UserServiceImpl) Create(ctx context.Context, request *dto.CreateUserRequest) (*domain.User, error) {
	var (
		hash string
		err  error
	)

	if hash, err = u.passwordSrv.Generate(request.Password); err != nil {
		return nil, err
	}

	request.Password = hash

	return u.repository.Create(ctx, request)
}

func (u *UserServiceImpl) FindOne(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return u.repository.FindOne(ctx, id)
}

func (u *UserServiceImpl) DeleteOne(ctx context.Context, id uuid.UUID) error {
	return u.repository.DeleteOne(ctx, id)
}

func NewUserService(repository UserRepository, passwordSrv core.PasswordService) UserService {
	return &UserServiceImpl{
		repository:  repository,
		passwordSrv: passwordSrv,
	}
}
