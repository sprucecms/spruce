package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/sprucecms/spruce/api"
	"github.com/sprucecms/spruce/cms"
	"github.com/sprucecms/spruce/sprucelib"
)

const apiPrefix = "/api/v1"  // TODO /api/ part should be a config param (and it should be passed dynamically to the JS
const adminPrefix = "/admin" // TODO /admin/ part should be a config param

func main() {
	app := sprucelib.NewSpruceApp()

	apiHandler := api.MountAt(apiPrefix, app)
	cmsHandler := cms.MountAt("/", app)
	adminHandler := http.StripPrefix(adminPrefix, http.FileServer(http.Dir(app.AdminDir())))

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
