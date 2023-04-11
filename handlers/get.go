package handlers

import (
	"encoding/json"
	"go/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func Get(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do if: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao atualizar objeto %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(todo)

}
