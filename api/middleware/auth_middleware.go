package middleware

import (
	"context"
	"go_template/api/schema"
	"go_template/domain/entity"
	"go_template/domain/entity_const"
	"go_template/log"
	"go_template/usecase/input_port"
	"strings"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

var (
	ErrNoAuthorizationHeader = entity_const.NewValidationErrorFromMsg("no authorization header passed")
	ErrNotSystemAdmin        = entity_const.NewValidationErrorFromMsg("you are not system admin")
)

type AuthMiddleware struct {
	UserUC input_port.IUserUseCase
}

func NewAuthMiddleware(userUC input_port.IUserUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		UserUC: userUC,
	}
}

// Authenticate
// tokenを取得して、認証するmiddlewareの例
func (m *AuthMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	logger, _ := log.NewLogger()

	return func(c echo.Context) error {
		// Get JWT Token From Header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, schema.TokenType+" ") {
			logger.Warn("Failed to authenticate", zap.Error(ErrNoAuthorizationHeader))
			return entity_const.NewUnauthorizedErrorFromMsg("token is invalid")
		}
		token := strings.TrimPrefix(authHeader, schema.TokenType+" ")

		// Authenticate
		userID, err := m.UserUC.Authenticate(token)
		if err != nil {
			logger.Warn("Failed to authenticate", zap.Error(err))
			return entity_const.NewUnauthorizedErrorFromMsg("token is invalid")
		}

		// set user detail to context
		user, err := m.UserUC.FindByID(userID)
		if err != nil {
			logger.Warn("Failed to find me", zap.Error(err))
			return entity_const.NewUnauthorizedErrorFromMsg("token is invalid")
		}
		c = SetToContext(c, user)

		return next(c)
	}
}

func SetToContext(c echo.Context, user entity.User) echo.Context {
	ctx := c.Request().Context()
	ctx = SetUserToContext(ctx, user)
	c.SetRequest(c.Request().WithContext(ctx))
	return c
}

type ContextKey string

var (
	userKey ContextKey = "userKey"
)

func SetUserToContext(ctx context.Context, user entity.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}