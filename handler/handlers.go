package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"solidgate-test/model"
	"solidgate-test/service"
	"solidgate-test/util"
)

type ValidatorHandlerInterface interface {
	PostCard(c *echo.Context) error
}

type ValidatorHandler struct {
	validatorService service.ValidatorInterface
}

func NewValidatorHandler(validator service.ValidatorInterface) *ValidatorHandler {
	return &ValidatorHandler{validatorService: validator}
}

func (vh *ValidatorHandler) ValidateCard(c echo.Context) error {
	var card model.Card

	if err := c.Bind(&card); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Validness: false,
			Error: util.ValidationError("Invalid request payload")})
	}

	if err := vh.validatorService.IsValidCard(card); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Validness: false, Error: err})
	}
	return c.JSON(http.StatusOK, model.Response{Validness: true})
}
