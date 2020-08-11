package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type inputPayload struct {
	ClientID     string `json:"client_id"`
	ClientSecret string
	Code         string `json:"code"`
	GrantType    string
	RedirectURI  string `json:"redirect_uri"`
}

type outputPayload struct {
	AccessToken  string `json:"access_token,omitempty"`
	UserID       int    `json:"user_id,omitempty"`
	ErrorType    string `json:"error_type,omitempty"`
	Code         int    `json:"code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func (api *API) handlePostAuthorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err)
				api.respond(w, r, http.StatusInternalServerError, err)
			}
		}()

		input := new(inputPayload)
		err = json.NewDecoder(r.Body).Decode(input)
		if err != nil {
			return
		}
		input.ClientSecret = api.cfg.InstagramSecret
		input.GrantType = "authorization_code"

		output, err := postInstagramRequest(input)
		if err != nil {
			return
		}
		api.respond(w, r, http.StatusOK, output)
	}
}

func postInstagramRequest(input *inputPayload) (*outputPayload, error) {
	target := "https://api.instagram.com/oauth/access_token"

	data := url.Values{}
	data.Set("client_id", input.ClientID)
	data.Set("client_secret", input.ClientSecret)
	data.Set("code", input.Code)
	data.Set("grant_type", input.GrantType)
	data.Set("redirect_uri", input.RedirectURI)
	req, _ := http.NewRequest("POST", target, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	var res *http.Response
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	output := new(outputPayload)
	err = json.NewDecoder(res.Body).Decode(output)
	return output, err
}
