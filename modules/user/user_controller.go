package user

import (
	"net/http"

	"github.com/alexkazantsev/templ-api/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	FindOne(context *gin.Context)
}

type UserControllerImpl struct {
	service UserService
}

func NewUserController(service UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}

func (u *UserControllerImpl) FindOne(ctx *gin.Context) {
	var (
		err  error
		id   uuid.UUID
		user *domain.User
	)

	if id, err = uuid.Parse(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid"})

		return
	}

	if user, err = u.service.FindOne(ctx, id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})

		return
	}

	ctx.JSON(http.StatusOK, user)
}
