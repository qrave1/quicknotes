package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/qrave1/logwrap"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/infrastructure/interfaces/http/dto"
	"net/http"
	"strconv"
)

type Note interface {
	HandleCreateNote(c echo.Context) error
	HandleReadNote(c echo.Context) error
	HandleReadAll(c echo.Context) error
	HandleUpdateNote(c echo.Context) error
	HandleDeleteNote(c echo.Context) error
}

type NoteController struct {
	noteUsecase domain.NoteUsecase
	log         logwrap.Logger
}

func NewNoteController(nu domain.NoteUsecase, log logwrap.Logger) *NoteController {
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

	id, err := n.noteUsecase.Create(ctx, note)
	if err != nil {
		n.log.Errorf("error create note. %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": id})
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

	return c.JSON(http.StatusOK, note)
}

func (n *NoteController) HandleReadAll(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.DefaultNoteRequest
	if err := c.Bind(&request); err != nil {
		n.log.Errorf("error bind create note request. %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	notes, err := n.noteUsecase.ReadAll(ctx, request.FolderId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, notes)
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
