package middlewares

import (
	"github.com/rs/cors"
)

func NewDefaultCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-Content-Type", "X-Serve-By"},
		AllowCredentials: true,
	})
}
