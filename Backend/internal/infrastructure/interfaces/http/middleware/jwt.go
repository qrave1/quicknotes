package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/qrave1/quicknotes/internal/usecase/auth"
	"net/http"
	"strconv"
)

func JwtMiddleware(secret []byte) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: secret,
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
		},
		SuccessHandler: func(c echo.Context) {
			user, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return
			}

			subj, err := user.Claims.(jwt.MapClaims).GetSubject()
			if err != nil {
				return
			}

			id, err := strconv.Atoi(subj)
			if err != nil {
				return
			}

			ctx := auth.SetUserId(c.Request().Context(), id)
			c.SetRequest(c.Request().WithContext(ctx))
		},
	})
}
