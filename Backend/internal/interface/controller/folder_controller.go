package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/qrave1/logwrap"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/infrastructure/interfaces/http/dto"
	"net/http"
	"strconv"
)

type Folder interface {
	HandleCreateFolder(c echo.Context) error
	HandleReadFolder(c echo.Context) error
	HandleReadFolders(c echo.Context) error
	HandleUpdateFolder(c echo.Context) error
	HandleDeleteFolder(c echo.Context) error
}

type FolderController struct {
	folderUsecase domain.FolderUsecase
	log           logwrap.Logger
}

func NewFolderController(fu domain.FolderUsecase, log logwrap.Logger) *FolderController {
	return &FolderController{folderUsecase: fu, log: log}
}

func (f *FolderController) HandleCreateFolder(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.CreateFolderRequest
	if err := c.Bind(&request); err != nil {
		f.log.Errorf("error bind create folder request. %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	folder := dto.FolderFromDTO(request)

	id, err := f.folderUsecase.Create(ctx, folder)
	if err != nil {
		f.log.Errorf("error create folder. %v", err)
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": id})
}

func (f *FolderController) HandleReadFolder(c echo.Context) error {
	ctx := c.Request().Context()

	v := c.Param("id")
	id, err := strconv.Atoi(v)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	folder, err := f.folderUsecase.FolderById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, folder)
}

func (f *FolderController) HandleReadFolders(c echo.Context) error {
	ctx := c.Request().Context()

	folders, err := f.folderUsecase.Folders(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, folders)
}

func (f *FolderController) HandleUpdateFolder(c echo.Context) error {
	ctx := c.Request().Context()

	v := c.Param("id")
	id, err := strconv.Atoi(v)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	var request dto.UpdateFolderRequest
	if err := c.Bind(&request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(request); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	folder := dto.FolderFromDTO(request)
	folder.Id = id

	if err := f.folderUsecase.Update(ctx, folder); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (f *FolderController) HandleDeleteFolder(c echo.Context) error {
	ctx := c.Request().Context()

	v := c.Param("id")
	id, err := strconv.Atoi(v)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err = f.folderUsecase.Delete(ctx, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
