// Copyright 2016 Aaron Longwell. All rights reserved.
// Use of this source code is governed by the Apache License
// that can be found in the LICENSE file.

/*
Package admin pertains to the Spruce administration area, the Javascript
web application which allows editing content in the CMS.
*/
package admin

import (
	"net/http"
	"regexp"

	"github.com/sprucecms/spruce/sprucelib"
)

// Returns an http.HandlerFunc that mounts the admin application at the
// supplied prefix.
//
// In order to support HTML5 History API functionality, URLs without a
// file extension suffix are internally redirected to the root URL (which
// will be served by index.html).
//
func MountAt(prefix string, app *sprucelib.SpruceApp) http.HandlerFunc {

	// Regexp of file extension paths
	filesWithExtensions := regexp.MustCompile("(?i)\\.(.{2,4})")

	// HTTP fileserver which will serve static files from the admin directory
	fileServer := http.StripPrefix(prefix, http.FileServer(http.Dir(app.AdminDir())))

	return func(w http.ResponseWriter, r *http.Request) {
		if !filesWithExtensions.MatchString(r.URL.Path) {
			// If the URL doesn't have a file extension (i.e. .png, .js, etc), then
			// it is likely a faux URL using the HTML5 History API, so we override
			// the URL to '/'
			r.URL.Path = prefix + "/" // TODO trailing double slash possibility?
		}
		fileServer.ServeHTTP(w, r)
	}
}
