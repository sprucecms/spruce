// Copyright 2016 Aaron Longwell. All rights reserved.
// Use of this source code is governed by the Apache License
// that can be found in the LICENSE file.

/*
Package cms pertains to the CMS front-end resolver. This code handles
all URIs that aren't handled by either the admin or API components.
*/
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

// Returns an http.Handler that mounts the CMS features at the
// supplied prefix.
//
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
