package middleware

import (
	"mail-api/common"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type middleware struct {
	ZapLogger *zap.Logger
}

func NewMiddleware(zapLogger *zap.Logger) *middleware {
	return &middleware{
		ZapLogger: zapLogger,
	}
}

func (m *middleware) JSONMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Accepts(common.ApplicationJSON)
		return c.Next()
	}
}

func (m *middleware) ContextLocaleMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(common.LocaleKey, c.Query(common.LocaleKey))
		return c.Next()
	}
}

func (m *middleware) LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Request().Header.Peek(common.XRequestID) == nil {
			c.Request().Header.Add(common.XRequestID, uuid.New().String())
		}

		logger := m.ZapLogger.With(zap.String(common.XRequestID, string(c.Request().Header.Peek(common.XRequestID))))

		logger.Debug(common.RequestInfoMsg,
			zap.String("method", string(c.Request().Header.Method())),
			zap.String("host", string(c.Request().Header.Host())),
			zap.String("path_uri", c.Request().URI().String()),
			zap.String("remote_addr", c.Context().RemoteAddr().String()),
			zap.String("body", string(c.Request().Body())),
		)

		if err := c.Next(); err != nil {
			return err
		}
		logger.Debug(common.ResponseInfoMsg,
			zap.String("body", string(c.Response().Body())),
		)
		return nil
	}
}
