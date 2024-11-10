package user

import (
	"context"

	"github.com/alexkazantsev/go-templ-api/domain"
	"github.com/alexkazantsev/go-templ-api/modules/core"
	"github.com/alexkazantsev/go-templ-api/modules/user/dto"
	"github.com/google/uuid"
)

type UserService interface {
	FindOne(context.Context, uuid.UUID) (*domain.User, error)
	Create(context.Context, *dto.CreateUserRequest) (*domain.User, error)
}

type UserServiceImpl struct {
	repository  UserRepository
	passwordSrv core.PasswordService
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

func NewUserService(repository UserRepository, passwordSrv core.PasswordService) UserService {
	return &UserServiceImpl{
		repository:  repository,
		passwordSrv: passwordSrv,
	}
}
