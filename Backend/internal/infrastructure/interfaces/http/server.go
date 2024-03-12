package http

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/qrave1/quicknotes/internal/config"
	"github.com/qrave1/quicknotes/internal/infrastructure/interfaces/http/middleware"
	"github.com/qrave1/quicknotes/internal/interface/controller"
	"github.com/qrave1/quicknotes/pkg/validator"
)

type Server struct {
	e *echo.Echo
}

func NewServer(
	cfg *config.Config,
	validate validator.Validator,
	authController controller.Auth,
	folderController controller.Folder,
	noteController controller.Note,
) *Server {
	s := &Server{e: echo.New()}
	s.e.Validator = validate
	s.e.Use(echomiddleware.Recover())

	jwt := middleware.JwtMiddleware([]byte(cfg.Server.Secret))

	auth := s.e.Group("")
	{
		auth.POST("/signup", authController.HandleSignUp)
		auth.POST("/signin", authController.HandleSignIn)
	}

	folders := s.e.Group("/folders", jwt)
	{
		folders.GET("/:id", folderController.HandleReadFolder)
		folders.POST("/", folderController.HandleCreateFolder)
		folders.PUT("/:id", folderController.HandleUpdateFolder)
		folders.DELETE("/:id", folderController.HandleDeleteFolder)
	}
}
