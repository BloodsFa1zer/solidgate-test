package router

import (
	"github.com/labstack/echo/v4"
	"solidgate-test/handler"
	"solidgate-test/service"
)

var validatorHandler = handler.NewValidatorHandler(service.NewCardValidator())

func UserRoute(e *echo.Echo) {

	e.POST("/card-validation", validatorHandler.ValidateCard)
}
