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

const apiPrefix = "/api/v1"
const adminPrefix = "/admin"

func main() {
	app := sprucelib.NewSpruceApp()

	apiHandler := api.MountAt(apiPrefix, app)
	cmsHandler := cms.MountAt("/", app)

	adminHandler := http.StripPrefix(adminPrefix, http.FileServer(http.Dir(app.AdminDir())))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("URL: '%s'\n", r.URL.Path)
		if strings.HasPrefix(r.URL.Path, apiPrefix) {
			fmt.Printf("apiHandler")
			apiHandler.ServeHTTP(w, r)
		} else if strings.HasPrefix(r.URL.Path, adminPrefix) {
			fmt.Printf("adminHandler")
			adminHandler.ServeHTTP(w, r)
		} else {
			fmt.Printf("cmsHandler")
			cmsHandler.ServeHTTP(w, r)
		}
	})

	// An error here indicates the server failed to start.
	// Exit ungracefully.
	httpErr := http.ListenAndServe(":1818", nil)
	if httpErr != nil {
		panic(httpErr)
		os.Exit(1)
	}

	// Wait for SIGINT or SIGTERM to cancel the process
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
