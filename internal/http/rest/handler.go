package rest

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"gitlab.com/HamelBarrer/game-server/internal/controller/role"
	"gitlab.com/HamelBarrer/game-server/internal/controller/user"
	midd "gitlab.com/HamelBarrer/game-server/internal/middleware"
)

func Handler() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/login", user.Login)
		r.Post("/", user.CreateUser)
		r.With(midd.VerificationToken).Get("/", user.ListUser)
		r.With(midd.VerificationToken).Get("/{id}", user.GetUser)
	})

	r.Route("/api/v1/roles", func(r chi.Router) {
		r.Get("/{id}", role.GetRole)
		r.Get("/", role.ListRole)
		r.Post("/", role.CreateRole)
		r.Put("/{id}", role.UpdateRole)
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	handler := cors.AllowAll().Handler(r)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
