package rest

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"gitlab.com/HamelBarrer/game-server/internal/controller/user"
)

func Handler() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/login", user.Login)
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	handler := cors.AllowAll().Handler(r)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
