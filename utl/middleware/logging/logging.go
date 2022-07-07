package logging

import (
	"go-alodokter/utl/middleware/request"

	"github.com/labstack/echo/v4"
)

func Logging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := request.Id()
			c.Set("request_id", reqID)
			// defer func(now time.Time) {

			// 	// flush cache
			// 	// cache.Delete(reqID)

			// 	message := tpl.LLvlAccess
			// 	fields := []zap.Field{
			// 		zap.String("at", now.Format(tpl.TimeFormat)),
			// 		zap.String("method", c.Request().Method),
			// 		zap.String("uri", c.Request().URL.String()),
			// 		zap.String("ip", c.RealIP()),
			// 		zap.String("host", c.Request().Host),
			// 		zap.String("user_agent", c.Request().UserAgent()),
			// 		zap.Int("code", c.Response().Status),
			// 	}
			// 	logger.WithRequestID(reqID).Info(message, fields...)
			// }(time.Now())
			return next(c)
		}
	}
}
