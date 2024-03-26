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

	// AUTH
	auth := s.e.Group("")
	auth.POST("/signup", authController.HandleSignUp)
	auth.POST("/signin", authController.HandleSignIn)

	// FOLDERS
	s.e.GET("/folders/:id", folderController.HandleReadFolder, jwt)
	s.e.GET("/folders", folderController.HandleReadFolders, jwt)
	s.e.POST("/folders", folderController.HandleCreateFolder, jwt)
	s.e.PUT("/folders/:id", folderController.HandleUpdateFolder, jwt)
	s.e.DELETE("/folders/:id", folderController.HandleDeleteFolder, jwt)

	// NOTES
	s.e.GET("/folders/:folder_id/notes/:id", noteController.HandleReadNote, jwt)
	s.e.GET("/folders/:folder_id/notes", noteController.HandleReadAll, jwt)
	s.e.POST("/folders/:folder_id/notes", noteController.HandleCreateNote, jwt)
	s.e.PUT("/folders/:folder_id/notes/:id", noteController.HandleUpdateNote, jwt)
	s.e.DELETE("/folders/:folder_id/notes/:id", noteController.HandleDeleteNote, jwt)

	return s
}

func (s *NoteServer) Run(cfg *config.Config) error {
	return s.e.Start(fmt.Sprintf(":%s", cfg.Server.Port))
}

func (s *NoteServer) Shutdown() {
	_ = s.e.Shutdown(context.Background())
}
