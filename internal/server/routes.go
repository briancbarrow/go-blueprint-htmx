package server

import (
	"net/http"

	"go-blueprint-htmx/cmd/api/github"
	"go-blueprint-htmx/cmd/web"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/css/*", fileServer)
	r.Handle("/js/*", fileServer)
	r.Get("/app", templ.Handler(web.Home()).ServeHTTP)
	r.Get("/auth", templ.Handler(web.AuthPage()).ServeHTTP)

	r.Get("/get_token", github.HandleGetToken)

	r.Post("/update_structure", web.UpdateStructureHandler)

	return r
}
