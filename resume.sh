#!/bin/bash

When you last left this project, you had determined that using Ember JS would be the 
best client-side system. It provides the most useful tooling out of the box, and is the
best when working with multiple developers.

You also determined that it was OK to use ember-simple-auth instead of rolling your
own. The simple auth system abstracts all the messy stuff and allows sufficient
customization. app/authenticators/* was updated to get tokens from /api/v1/token
and that URL was set up in go to return a hard-coded response. You evaluated the OSIN
system, but decided that system overreached... the requirement of a Storage system.

Right now the login is working.

open "http://brewhouse.io/blog/2015/02/12/ember-vs-angular-authentication.html"
open "https://github.com/BrewhouseTeam/ember-auth-example/blob/master/app/templates/public.hbs"
open "http://localhost:1818/admin/"

Jun 12 - You began transferring some logic from the derbyphoto project to use gorm and PostgreSQL as the
         implementation of the DataStore interface. Login checking was just about ready to be worked.

         Start by opening:

         scm.cdinet.net/derbyphoto/models.go
         scm.cdinet.net/derbyphoto/cmd/derbyphoto/main.go
         github.com/sprucecms/spruce/api/api.go
         github.com/sprucecms/spruce/sprucelib/user.go
         github.com/sprucecms/spruce/sprucelib/datastore-gorm.go
