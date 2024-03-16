package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/qrave1/quicknotes/internal/config"
	"github.com/qrave1/quicknotes/internal/infrastructure/interfaces/http/middleware"
	"github.com/qrave1/quicknotes/internal/interface/controller"
	"github.com/qrave1/quicknotes/pkg/validator"
)

type NoteServer struct {
	e *echo.Echo
}

func NewNoteServer(
	cfg *config.Config,
	validate validator.Validator,
	authController controller.Auth,
	folderController controller.Folder,
	noteController controller.Note,
) *NoteServer {
	s := &NoteServer{e: echo.New()}
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

	notes := s.e.Group("folders/:folder_id/notes", jwt)
	{
		notes.GET("/:id", noteController.HandleReadNote)
		notes.GET("/", noteController.HandleReadAll)
		notes.POST("/", noteController.HandleCreateNote)
		notes.PUT("/:id", noteController.HandleUpdateNote)
		notes.DELETE("/:id", noteController.HandleDeleteNote)
	}

	return s
}

func (s *NoteServer) Run(cfg *config.Config) error {
	return s.e.Start(fmt.Sprintf(":%s", cfg.Server.Port))
}

func (s *NoteServer) Shutdown() {
	_ = s.e.Shutdown(context.Background())
}
