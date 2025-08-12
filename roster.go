package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) DeleteArmy(w http.ResponseWriter, r *http.Request) {

	str := chi.URLParam(r, "id")
	id, err := uuid.Parse(str)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to parse id")
		return
	}

	err = cfg.db.DeleteArmy(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to delete army")
	}

	respondWithJSON(w, 200, "army successfully deleted")
}
