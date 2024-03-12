package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/qrave1/logger-wrapper/logrus"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/infrastructure/interfaces/http/dto"
	"net/http"
	"strconv"
)

type Note interface {
	HandleCreateNote(c echo.Context) error
	HandleReadNote(c echo.Context) error
	HandleUpdateNote(c echo.Context) error
	HandleDeleteNote(c echo.Context) error
}

type NoteController struct {
	noteUsecase domain.NoteUsecase
	log         logrus.Logger
}

func NewNoteController(nu domain.NoteUsecase, log logrus.Logger) *NoteController {
	return &NoteController{noteUsecase: nu, log: log}
}

func (n *NoteController) HandleCreateNote(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.CreateNoteRequest
	if err := c.Bind(&request); err != nil {
		n.log.Errorf("error bind create note request. %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	note := dto.NoteFromDTO(request)

	if err := n.noteUsecase.Create(ctx, note); err != nil {
		n.log.Errorf("error create folder. %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}

func (n *NoteController) HandleReadNote(c echo.Context) error {
	ctx := c.Request().Context()

	v := c.Param("id")
	id, err := strconv.Atoi(v)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	note, err := n.noteUsecase.Read(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"note": note,
	})
}

func (n *NoteController) HandleUpdateNote(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.UpdateNoteRequest
	if err := c.Bind(&request); err != nil {
		n.log.Errorf("error bind create note request. %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	note := dto.NoteFromDTO(request)

	if err := n.noteUsecase.Update(ctx, note); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (n *NoteController) HandleDeleteNote(c echo.Context) error {
	ctx := c.Request().Context()

	v := c.Param("id")
	id, err := strconv.Atoi(v)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := n.noteUsecase.Delete(ctx, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
