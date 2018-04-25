package server

import (
	"net/http"
	"os"

	"github.com/gorilla/pat"
	"github.com/rs/cors"
)

var (
	authorsPath  = "/authors"
	commentsPath = "/comments"
)

func InitServer() {
	p := pat.New()

	p.Get(authorsPath, CreateAuthor)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-API-KEY", "Origin", "X-Requested-With", "X-Session", "Content-Type", "Accept", "Access-Control-Request-Method"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(p)

	if err := http.ListenAndServe(os.Getenv("MICRO_MUX_ADDRESS"), handler); err != nil {
		panic(err)
	}
}
