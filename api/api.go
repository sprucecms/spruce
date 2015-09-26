package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sprucecms/spruce/sprucelib"
)

const Version string = "0.0.1"

type apiManager struct {
	prefix string
	router *mux.Router
	app    *sprucelib.SpruceApp
}

func MountAt(prefix string, app *sprucelib.SpruceApp) http.Handler {
	m := apiManager{prefix: prefix}
	m.router = mux.NewRouter().PathPrefix(m.prefix).Subrouter()
	m.app = app

	m.router.HandleFunc("/", m.apiMetadata)

	return m.router
}

// GET /<APIPrefix>/
//
// Returns version information for the API
//
func (m apiManager) apiMetadata(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Version string
	}{Version: Version}
	json.NewEncoder(w).Encode(data)
}
