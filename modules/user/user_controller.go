package user

import (
	"fmt"

	"github.com/alexkazantsev/go-templ-api/domain"
	"github.com/alexkazantsev/go-templ-api/modules/user/dto"
	"github.com/alexkazantsev/go-templ-api/pkg/xcall"
	"github.com/alexkazantsev/go-templ-api/pkg/xerror"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	service UserService
}

// UpdateOne /users/{id}
func (u *UserController) UpdateOne(ctx *gin.Context) {
	status, response := xcall.CallM(func() (*domain.User, error) {
		var (
			id      uuid.UUID
			request dto.UpdateUserRequest
			err     error
		)

		if id, err = uuid.Parse(ctx.Param("id")); err != nil {
			return nil, fmt.Errorf("id must be a valid uuid: %w", xerror.ErrInvalidRequest)
		}

		if err = ctx.ShouldBindJSON(&request); err != nil {
			return nil, fmt.Errorf("binding error: %w: %w", err, xerror.ErrInvalidRequest)
		}

		request.ID = id

		if err = request.Validate(); err != nil {
			return nil, err
		}

		return u.service.UpdateOne(ctx, &request)
	})

	ctx.JSON(status, response)
}

func (u *UserController) Create(ctx *gin.Context) {
	status, response := xcall.CallM(func() (*domain.User, error) {
		var (
			user    *domain.User
			request dto.CreateUserRequest
			err     error
		)

		if err = ctx.ShouldBindJSON(&request); err != nil {
			return nil, fmt.Errorf("binding error: %w: %w", err, xerror.ErrInvalidRequest)
		}

		if err = request.Validate(); err != nil {
			return nil, err
		}

		if user, err = u.service.Create(ctx, &request); err != nil {
			return nil, err
		}

		return user, nil
	})

	ctx.JSON(status, response)
}

func NewUserController(service UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (u *UserController) FindOne(ctx *gin.Context) {
	var (
		err  error
		id   uuid.UUID
		user *domain.User
	)

	status, payload := xcall.CallM[*domain.User](func() (*domain.User, error) {
		if id, err = uuid.Parse(ctx.Param("id")); err != nil {
			return nil, fmt.Errorf("id must be a valid uuid: %w", xerror.ErrInvalidRequest)
		}

		if user, err = u.service.FindOne(ctx, id); err != nil {
			return nil, err
		}

		return user, nil
	})

	ctx.JSON(status, payload)
}

func (u *UserController) FindMany(ctx *gin.Context) {
	status, payload := xcall.CallM(func() ([]*domain.User, error) {
		var (
			err error
			req dto.FindManyRequest
		)

		if err = ctx.ShouldBindQuery(&req); err != nil {
			return nil, xerror.ErrInvalidRequest
		}

		if err = req.Validate(); err != nil {
			return nil, err
		}

		return u.service.FindMany(ctx, &req)
	})

	ctx.JSON(status, payload)
}

func (u *UserController) DeleteOne(ctx *gin.Context) {
	status, payload := xcall.CallS(func() error {
		var (
			err error
			id  uuid.UUID
		)

		if id, err = uuid.Parse(ctx.Param("id")); err != nil {
			return fmt.Errorf("id must be a valid uuid: %w", xerror.ErrInvalidRequest)
		}

		return u.service.DeleteOne(ctx, id)
	})

	ctx.JSON(status, payload)
}
