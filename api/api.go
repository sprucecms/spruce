// Copyright 2016 Aaron Longwell. All rights reserved.
// Use of this source code is governed by the Apache License
// that can be found in the LICENSE file.

/*
Package api pertains to the Spruce API server-side components. All API
handlers are implemented in this package.
*/
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

// Returns an http.Handler which mounts the API URI handlers at the supplied
// prefix.
//
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
