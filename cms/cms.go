package cms

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	lib "github.com/sprucecms/spruce/sprucelib"
)

type cmsManager struct {
	prefix string
	router *mux.Router
	app    *lib.SpruceApp
}

func MountAt(prefix string, app *lib.SpruceApp) http.Handler {
	m := cmsManager{prefix: prefix}
	m.router = mux.NewRouter()
	if m.prefix != "/" && m.prefix != "" {
		m.router = m.router.PathPrefix(m.prefix).Subrouter()
	}

	m.router.NotFoundHandler = http.HandlerFunc(m.handler)

	return m.router
}

func (m cmsManager) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "No rendering features enabled yet. Path: '%s'", r.URL.Path)
}
