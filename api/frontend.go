package main

import "net/http"

func (api *API) handleGetFrontend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, api.cfg.FrontendApp, http.StatusSeeOther)
	}
}
