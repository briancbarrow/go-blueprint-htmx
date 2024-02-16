package server

import (
	"encoding/json"
	"log"
	"net/http"

	"go-blueprint-htmx/cmd/web"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/css/*", fileServer)
	r.Handle("/js/*", fileServer)
	r.Get("/web", templ.Handler(web.Home()).ServeHTTP)
	r.Post("/hello", web.HelloWebHandler)

	r.Post("/update_structure", web.UpdateStructureHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
