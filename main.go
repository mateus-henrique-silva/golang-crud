package main

import (
	"fmt"
	"go/configs"
	"go/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()

	if err != nil {
		log.Fatalf("Error ao carregar configuracoes no programa", err.Error())
	}
	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Updated)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)

	port := configs.GetServerPort()
	log.Printf("Iniciando servidor na porta %s", port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)

	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %s", err.Error())
	}
}
