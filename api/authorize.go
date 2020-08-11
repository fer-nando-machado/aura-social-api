package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (api *API) handlePostAuthorize() http.HandlerFunc {
	type payload struct {
		ClientID     string `json:"client_id"`
		ClientSecret string
		Code         string `json:"code"`
		GrantType    string `json:"grant_type"`
		RedirectURI  string `json:"redirect_uri"`
	}

	type outputPayload struct {
		AccessToken  string `json:"access_token,omitempty"`
		UserID       string `json:"user_id,omitempty"`
		ErrorType    string `json:"error_type,omitempty"`
		Code         int    `json:"code,omitempty"`
		ErrorMessage string `json:"error_message,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err)
				api.respond(w, r, http.StatusInternalServerError, "Error.")
			}
		}()

		// PARSE INPUT
		input := new(payload)
		err = json.NewDecoder(r.Body).Decode(input)
		if err != nil {
			return
		}

		// MAKE REQUEST
		resp, err := http.Post("https://api.instagram.com/oauth/access_token", "application/json", r.Body)
		if err != nil {
			return
		}
		// TODO do something wish status code?
		log.Println(resp.StatusCode)

		defer resp.Body.Close()
		// instaBody, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	return
		// }
		result := new(outputPayload)
		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return
		}

		// RETURN PAYLOAD
		// output := payload{
		// 	ClientID:     input.ClientID,
		// 	ClientSecret: api.cfg.InstagramSecret,
		// 	Code:         input.Code,
		// 	GrantType:    input.GrantType,
		// 	RedirectURI:  input.RedirectURI,
		// }
		api.respond(w, r, http.StatusOK, result)
	}
}
