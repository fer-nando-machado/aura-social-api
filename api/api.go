package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

type API struct {
	cfg *Config
}

func (api *API) Routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(cors.New(cors.Options{
		AllowedOrigins: []string{api.cfg.AllowedOrigin},
		Debug:          true,
	}).Handler)

	mux.Get("/", api.handleGetFrontend())
	mux.Get("/healthcheck", api.handleGetHealthcheck())
	mux.Post("/authorize", api.handlePostAuthorize())

	return mux
}

func (api *API) respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if data == nil {
		return
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(buffer.Bytes())
}
