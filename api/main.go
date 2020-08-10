package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := NewConfig()

	api := API{cfg: cfg}
	log.Println("Initalizing server at port " + cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, api.Routes()); err != nil {
		log.Fatalf("Failed to initialize server: %s", err)
	}
}
