package middleware

import (
	"errors"
	"go_template/domain/entity_const"
	"go_template/log"
	"net/http"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

type ErrorMiddleware struct {
}

func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{}
}

func (m *ErrorMiddleware) HandleError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		logger, _ := log.NewLogger()
		logger.Info("failed API", zap.Error(err))
		switch {
		case errors.As(err, new(*entity_const.ValidationError)):
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		case errors.As(err, new(*entity_const.UnauthorizedError)):
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		case errors.As(err, new(*entity_const.HasNoPermissionError)):
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		case errors.As(err, new(*entity_const.NotFoundError)):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case errors.As(err, new(*entity_const.ConflictError)):
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		case errors.As(err, new(*entity_const.DatabaseError)):
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
}
