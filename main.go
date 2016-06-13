package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sprucecms/spruce/admin"
	"github.com/sprucecms/spruce/api"
	"github.com/sprucecms/spruce/cms"
	"github.com/sprucecms/spruce/sprucelib"
)

const apiPrefix = "/api/v1"  // TODO /api/ part should be a config param (and it should be passed dynamically to the JS)
const adminPrefix = "/admin" // TODO /admin/ part should be a config param to support security-through-obscurity strategies

func main() {

	db, err := gorm.Open("postgres", "host=localhost user=root dbname=dev_spruce sslmode=disable password=")
	db.AutoMigrate(&sprucelib.User{})
	if err != nil {
		panic(err)
	}
	app := sprucelib.NewSpruceApp()
	app.DataStore = sprucelib.SqlDataStore{DB: db}

	apiHandler := api.MountAt(apiPrefix, app)
	cmsHandler := cms.MountAt("/", app)
	adminHandler := admin.MountAt(adminPrefix, app)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, apiPrefix) {
			apiHandler.ServeHTTP(w, r)
		} else if strings.HasPrefix(r.URL.Path, adminPrefix) {
			adminHandler.ServeHTTP(w, r)
		} else {
			cmsHandler.ServeHTTP(w, r)
		}
	})

	// An error here indicates the server failed to start.
	// Exit ungracefully.
	go (func() {
		httpErr := http.ListenAndServe(":1818", nil) // TODO Port should be config param
		if httpErr != nil {
			panic(httpErr)
			os.Exit(1)
		}
	})()

	// Wait for SIGINT or SIGTERM to cancel the process
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	r := <-ch
	fmt.Printf("Spruce shut down: %s", r)
}
