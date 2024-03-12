package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/qrave1/logger-wrapper/logrus"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/infrastructure/interfaces/http/dto"
	"net/http"
)

type Auth interface {
	HandleSignUp(c echo.Context) error
	HandleSignIn(c echo.Context) error
}

type AuthController struct {
	userUsecase domain.UserUsecase
	validator   validator.Validate
	log         logrus.Logger
}

func NewAuthController(uu domain.UserUsecase, log logrus.Logger) *AuthController {
	return &AuthController{userUsecase: uu, log: log}
}

func (a *AuthController) HandleSignUp(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.SignUpRequest
	if err := c.Bind(&request); err != nil {
		a.log.Errorf("error bind signUpRequest. %v", err)
		return c.NoContent(400)
	}

	if err := a.validator.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user := dto.UserFromDTO(&request)

	if err := a.userUsecase.SignUp(ctx, user); err != nil {
		a.log.Errorf("error create new user. %v", err)
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *AuthController) HandleSignIn(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.SignInRequest
	if err := c.Bind(&request); err != nil {
		a.log.Errorf("error bind signInRequest. %v", err)
		return c.NoContent(400)
	}

	if err := a.validator.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user := dto.UserFromDTO(&request)

	token, err := a.userUsecase.SignIn(ctx, user)
	if err != nil {
		a.log.Errorf("error create new user. %v", err)
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
