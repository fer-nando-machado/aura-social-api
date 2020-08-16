package main

import "net/http"

func (api *API) handleGetHealthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(api.cfg.Port))
	}
}
