package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/qrave1/logger-wrapper/logrus"
	"github.com/qrave1/quicknotes/internal/domain"
)

type Note interface {
	HandleCreateNote(c echo.Context) error
	HandleReadNote(c echo.Context) error
	HandleUpdateNote(c echo.Context) error
	HandleDeleteNote(c echo.Context) error
}

type NoteController struct {
	nu  domain.NoteUsecase
	log logrus.Logger
}

func (n *NoteController) HandleCreateNote(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (n *NoteController) HandleReadNote(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (n *NoteController) HandleUpdateNote(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (n *NoteController) HandleDeleteNote(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewNoteController(nu domain.NoteUsecase, log logrus.Logger) *NoteController {
	return &NoteController{nu: nu, log: log}
}
