package sprucelib

import (
	"fmt"
	"path/filepath"
)

type SpruceApp struct {
	adminDir   string
	DataStore  DataStore
	AssetStore AssetStore
}

type DataStore interface {
}

type AssetStore interface {
}

func NewSpruceApp() *SpruceApp {
	return &SpruceApp{}
}

func (app *SpruceApp) AdminDir() string {
	dir := "./web/admin"
	if app.adminDir != "" {
		dir = app.adminDir
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Admin Dir: '%s'\n", dir)
	return dir
}
