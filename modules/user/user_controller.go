package user

import (
	"net/http"

	"github.com/alexkazantsev/go-templ-api/domain"
	"github.com/alexkazantsev/go-templ-api/modules/user/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	FindOne(*gin.Context)
	Create(*gin.Context)
}

type UserControllerImpl struct {
	service UserService
}

func (u *UserControllerImpl) Create(ctx *gin.Context) {
	var (
		user    *domain.User
		request dto.CreateUserRequest
		err     error
	)

	if err = ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err = request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if user, err = u.service.Create(ctx, &request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, user)
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
