package sprucelib

import (
	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

type SpruceApp struct {
	adminDir   string
	Logger     *log.Logger
	DataStore  DataStore
	AssetStore AssetStore
}

type DataStore interface {
}

type AssetStore interface {
}

func NewSpruceApp() *SpruceApp {
	a := &SpruceApp{}
	a.Logger = log.New()
	return a
}

func (app *SpruceApp) AdminDir() string {
	dir := "./web/spruceadmin/dist"
	if app.adminDir != "" {
		dir = app.adminDir
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	app.Logger.WithFields(log.Fields{"adminDir": dir}).Info("CMS admin app path")
	return dir
}
