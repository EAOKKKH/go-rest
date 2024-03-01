package handler

import (
	"go-rest/internal/config"
	"go-rest/internal/models"
	"go-rest/internal/models/validators"
	"go-rest/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type IUrlHandler interface {
	CreateUrl(ctx echo.Context) error
	Redirect(ctx echo.Context) error
}

type urlHandler struct {
	cfg          *config.Config
	logger       *logrus.Logger
	urlUsecase   usecase.IUrlUsecase
	urlValidator validators.IUrlValidator
}

func NewUrlHandler(cfg *config.Config, logger *logrus.Logger, urlUseCase usecase.IUrlUsecase, urlValidator validators.IUrlValidator) IUrlHandler {
	return &urlHandler{
		cfg:          cfg,
		logger:       logger,
		urlUsecase:   urlUseCase,
		urlValidator: urlValidator,
	}
}

func (h *urlHandler) CreateUrl(ctx echo.Context) error {
	url := &models.Url{}
	if err := ctx.Bind(url); err != nil {
		h.logger.Errorf("Error binding from request body: %v", err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := h.urlValidator.Validate(url); err != nil {
		h.logger.Error(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	createdUrl, err := h.urlUsecase.CreateUrl(ctx.Request().Context(), url)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, createdUrl)

}

func (h *urlHandler) Redirect(ctx echo.Context) error {
	alias := ctx.Param("alias")
	url, err := h.urlUsecase.GetUrl(ctx.Request().Context(), alias)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.Redirect(http.StatusFound, url)
}

func MapUrlRoutes(group *echo.Group, h IUrlHandler) {
	group.POST("", h.CreateUrl)
	group.GET("/:alias", h.Redirect)
}
