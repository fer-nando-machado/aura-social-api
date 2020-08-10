package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

type API struct {
	cfg *Config
}

func (api *API) Routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/healthcheck", api.handleHealthCheck())
	return mux
}

func (api *API) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(api.cfg.Port))
	}
}
