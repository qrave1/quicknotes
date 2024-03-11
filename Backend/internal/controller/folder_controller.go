package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/qrave1/logger-wrapper/logrus"
	"github.com/qrave1/quicknotes/internal/domain"
)

type Folder interface {
	HandleCreateFolder(c echo.Context) error
	HandleReadFolder(c echo.Context) error
	HandleUpdateFolder(c echo.Context) error
	HandleDeleteFolder(c echo.Context) error
}

type FolderController struct {
	fu  domain.FolderUsecase
	log logrus.Logger
}

func NewFolderController(fu domain.FolderUsecase, log logrus.Logger) *FolderController {
	return &FolderController{fu: fu, log: log}
}

func (f *FolderController) HandleCreateFolder(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f *FolderController) HandleReadFolder(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f *FolderController) HandleUpdateFolder(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f *FolderController) HandleDeleteFolder(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

//func (f *FolderController) HandleCreateFolder(c echo.Context) error {
//	ctx := context.Background()
//
//	var request domain.Folder
//	if err := c.Bind(request); err != nil {
//		f.log.Warnf("deserialize folder error. %v", err)
//		return err
//	}
//
//	request.UserId = c.Get()
//	f.fu.SignIn(ctx)
//}
