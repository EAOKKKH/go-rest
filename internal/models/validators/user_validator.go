package validators

import (
	"go-rest/internal/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IUrlValidator interface {
	Validate(url *models.Url) error
}

type urlValidator struct {
}

func NewUrlValidator() IUrlValidator {
	return &urlValidator{}
}

func (v *urlValidator) Validate(url *models.Url) error {
	return validation.ValidateStruct(url,
		validation.Field(
			&url.Url,
			validation.Required.Error("url can't be empty"),
		),
	)
}
