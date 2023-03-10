package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jeswyrne/pig-latin/controller"
	"github.com/Jeswyrne/pig-latin/middlewares"
	"github.com/Jeswyrne/pig-latin/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const port = 3000

func main() {

	repo := repository.NewDatabase()
	contrl := controller.NewPigLatin(repo)

	r := chi.NewMux()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	// Custom Middleware
	r.Use(middlewares.SetMiddlewareHeaders)

	r.Get("/", contrl.GetHandler)
	r.Post("/", contrl.PostHandler)

	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 127.0.0.1:%v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
