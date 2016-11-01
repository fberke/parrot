package web

import (
	"github.com/anthonynsimon/parrot/datastore"
	"github.com/anthonynsimon/parrot/paths"
	"github.com/pressly/chi"
)

var store datastore.Store

func Register(router *chi.Mux, ds datastore.Store) {
	store = ds
	registerRoutes(router)
}

func registerRoutes(router *chi.Mux) {
	router.Get(paths.PingPath, webHandlerFunc(ping).ServeHTTP)
	router.Get("/login", webHandlerFunc(loginForm).ServeHTTP)
	router.Post("/login", webHandlerFunc(login).ServeHTTP)
	router.Get("/register", webHandlerFunc(newUser).ServeHTTP)
	router.Post("/register", webHandlerFunc(createUser).ServeHTTP)

	router.Route(paths.ProjectsPath, func(r chi.Router) {
		r.Get("/:projectID", webHandlerFunc(showProject).ServeHTTP)
		r.Get("/new", webHandlerFunc(newProject).ServeHTTP)

		r.Route("/:projectID"+paths.DocumentsPath, func(r chi.Router) {
			r.Get("/", webHandlerFunc(findDocuments).ServeHTTP)
			r.Get("/:documentID", webHandlerFunc(showDocument).ServeHTTP)
			r.Get("/new", webHandlerFunc(newDocument).ServeHTTP)
		})
	})

	router.Route(paths.UsersPath, func(r chi.Router) {
		r.Get("/:userID", webHandlerFunc(showUser).ServeHTTP)
	})
}