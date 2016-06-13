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
	"fmt"
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
	m.router.HandleFunc("/token", m.oauth2Token)

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

func (m apiManager) oauth2Token(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, err := m.app.DataStore.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("user.username", user.Username)

	// TODO actually authenticate
	data := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}{AccessToken: "TODOMAKETHISAREALTOKEN", TokenType: "bearer", ExpiresIn: 3153600000}
	json.NewEncoder(w).Encode(data)
}
