package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (cfg *ApiConfig) GetAbilities(w http.ResponseWriter, r *http.Request) {
	abilities, err := cfg.Db.GetAbilities(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to get abilities")
		return
	}
	abiltiesSlice := []Abilities{}
	for _, item := range abilities {
		abilitiesJSON := Abilities{
			DatasheetID: item.DatasheetID,
			Line:        item.Line,
			AbilityID:   item.AbilityID,
			Model:       item.Model,
			Name:        item.Name,
			Description: item.Description,
			Type:        item.Type,
			Parameter:   item.Parameter,
		}
		abiltiesSlice = append(abiltiesSlice, abilitiesJSON)
	}
	respondWithJSON(w, 200, abiltiesSlice)
}

func (cfg *ApiConfig) GetAbilitiesForModel(w http.ResponseWriter, r *http.Request) {

	Id := chi.URLParam(r, "id")

	parsedID, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to parse Id")
		return
	}

	Id32 := int32(parsedID)

	models, err := cfg.Db.GetAbilitiesForModel(r.Context(), Id32)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to fetch abilities")
		return
	}

	respondWithJSON(w, 200, models)
}
