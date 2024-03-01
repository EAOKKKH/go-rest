package handler

import (
	"go-rest/internal/models"
	"go-rest/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type IAuthHandler interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	Logout(ctx echo.Context) error
}

type authHandler struct {
	logger      *logrus.Logger
	authUsecase usecase.IAuthUsecase
}

func NewAuthHandler(logger *logrus.Logger, authUsecase usecase.IAuthUsecase) IAuthHandler {
	return &authHandler{
		logger:      logger,
		authUsecase: authUsecase,
	}
}

// @Summary login
// @Description returns user and set session-cookie
// @Tags auth
// @Accept json
// @Param input body models.User true "user params"
// @Produce json
// @Success 200 {object} models.User
// @Router /auth/signin [post]
func (h *authHandler) Login(ctx echo.Context) error {
	panic("unimplemented")
}

// Logout implements IAuthHandler.
func (h *authHandler) Logout(ctx echo.Context) error {
	panic("unimplemented")
}

// @Summary register new user
// @Description register new user, returns user
// @Tags auth
// @Accept json
// @Param input body models.User true "user params"
// @Produce json
// @Success 201 {object} models.User
// @Router /auth/signup [post]
func (h *authHandler) Register(ctx echo.Context) error {
	user := &models.User{}
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	createdUser, err := h.authUsecase.Register(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, createdUser)
}

func MapAuthRoutes(group *echo.Group, h IAuthHandler) {
	group.POST("/signup", h.Register)
	group.POST("/signin", h.Login)
}
