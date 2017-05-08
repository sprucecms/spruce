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

	jwt "github.com/dgrijalva/jwt-go"
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
	m.router.HandleFunc("/oauth2/token", m.oauth2Token)

	return m.router
}

func (m apiManager) requireAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
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

type tokenResponse struct {
	Error       string `json:"error,omitempty"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// POST /<APIPrefix>/oauth2/token
//
// Used to request an OAuth2 bearer token.
//
func (m apiManager) oauth2Token(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	fmt.Println(username)
	fmt.Println(password)

	user, err := m.app.DataStore.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(tokenResponse{Error: "invalid_client"})
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":      fmt.Sprintf("%d", user.ID),
			"username": fmt.Sprintf("%s", user.Username),
			"admin":    "true",
		})
		signedToken, err := token.SignedString([]byte("B8F24148A16532328E4EE5DC8DEAE")) // TODO Change this secret
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(tokenResponse{AccessToken: signedToken, TokenType: "bearer", ExpiresIn: 315600000})
	}
}
