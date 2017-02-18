package sprucelib

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type SpruceApp struct {
	adminDir   string
	Logger     *logrus.Logger
	DataStore  DataStore
	AssetStore AssetStore
}

type AssetStore interface {
}

func NewSpruceApp() *SpruceApp {
	a := &SpruceApp{}
	a.Logger = logrus.New()
	return a
}

func (app *SpruceApp) AdminDir() string {
	dir := "./web/admin/dist"
	if app.adminDir != "" {
		dir = app.adminDir
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	app.Logger.WithFields(logrus.Fields{"adminDir": dir}).Info("CMS admin app path")
	return dir
}
