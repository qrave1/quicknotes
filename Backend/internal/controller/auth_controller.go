package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/qrave1/logger-wrapper/logrus"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/usecase/auth"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Auth interface {
	HandleSignUp(c echo.Context) error
	HandleSignIn(c echo.Context) error
}

type AuthController struct {
	userUsecase domain.UserUsecase
	auth        auth.Auth
	log         logrus.Logger
}

func NewAuthController(uu domain.UserUsecase, log logrus.Logger) *AuthController {
	return &AuthController{userUsecase: uu, log: log}
}

func (a *AuthController) HandleSignUp(c echo.Context) error {
	ctx := context.Background()

	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		a.log.Errorf("error generate hashed password. %v", err)
		return err
	}

	u := domain.User{
		Name:     username,
		Email:    email,
		Password: string(passHash),
	}

	if err = a.userUsecase.Create(ctx, u); err != nil {
		a.log.Errorf("error create new user. %v", err)
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *AuthController) HandleSignIn(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.FormValue("email")
	password := c.FormValue("password")

	u, err := a.userUsecase.ReadByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		a.log.Infof("invalid credentials. user_id=%d", u.Id)
		return err
	}

	t, err := a.auth.Generate(u.Id)
	if err != nil {
		a.log.Warnf("error create jwt token. %v", err)
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
