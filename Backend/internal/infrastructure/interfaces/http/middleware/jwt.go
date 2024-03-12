package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/qrave1/quicknotes/internal/usecase/auth"
	"net/http"
	"strconv"
)

const (
	authorizationHeader = "X-Auth-Token"
)

func JwtMiddleware(secret []byte) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		//BeforeFunc: func(e echo.Context) {
		//	e.Request().Header.Set(authorizationHeader, "Bearer "+e.Request().Header.Get(authorizationHeader))
		//},
		SigningKey:  secret,
		TokenLookup: "header:" + authorizationHeader,
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
		},
		SuccessHandler: func(c echo.Context) {
			user, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return
			}

			claims, ok := user.Claims.(*jwt.RegisteredClaims)
			if !ok {
				return
			}

			id, err := strconv.Atoi(claims.Subject)
			if err != nil {
				return
			}

			ctx := auth.SetUserId(c.Request().Context(), id)
			c.SetRequest(c.Request().WithContext(ctx))
			//c.Set(CurrentUserKey, claims.Subject)
		},
	})
}
