package server

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/rs/cors"
	"github.com/saromanov/minimal/api"
)

var (
	authorsPath  = "/authors"
	commentsPath = "/comments"
)

// InitServer provides init of minimal server
func InitServer() {
	p := pat.New()
	p.Post(authorsPath, api.CreateAuthor)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-API-KEY", "Origin", "X-Requested-With", "X-Session", "Content-Type", "Accept", "Access-Control-Request-Method"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(p)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
