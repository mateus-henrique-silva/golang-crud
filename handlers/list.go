package handlers

import (
	"encoding/json"
	"go/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registro:%v", err)
	}
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
