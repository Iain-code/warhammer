package handlers

import (
	"encoding/json"
	"net/http"
	"warhammer/internal/db"

	"github.com/go-chi/chi/v5"
)

func (cfg *ApiConfig) GetEnhancementsForFaction(w http.ResponseWriter, r *http.Request) {
	Id := chi.URLParam(r, "id")
	enchancements, err := cfg.Db.GetEnhancementsForFaction(r.Context(), Id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to get enhancements")
		return
	}

	enhanceSlice := []Enhancement{}
	for _, item := range enchancements {
		enhanceJSON := Enhancement{
			ID:          item.ID,
			FactionID:   item.FactionID,
			Name:        item.Name,
			Cost:        item.Cost,
			Detachment:  item.Detachment,
			Description: item.Description,
		}
		enhanceSlice = append(enhanceSlice, enhanceJSON)
	}

	respondWithJSON(w, 200, enhanceSlice)

}

func (cfg *ApiConfig) GetEnhancements(w http.ResponseWriter, r *http.Request) {

	enhance, err := cfg.Db.GetEnhancements(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to get enhancements")
		return
	}

	enhanceSlice := []Enhancement{}
	for _, item := range enhance {
		enhanceJSON := Enhancement{
			ID:          item.ID,
			FactionID:   item.FactionID,
			Name:        item.Name,
			Cost:        item.Cost,
			Detachment:  item.Detachment,
			Description: item.Description,
		}
		enhanceSlice = append(enhanceSlice, enhanceJSON)
	}

	respondWithJSON(w, 200, enhanceSlice)
}

func (cfg *ApiConfig) AddNewEnhancement(w http.ResponseWriter, r *http.Request) {

	enhancement := Enhancement{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&enhancement)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	params := db.AddNewEnhancementParams{
		FactionID:   enhancement.FactionID,
		Name:        enhancement.Name,
		Cost:        enhancement.Cost,
		Detachment:  enhancement.Detachment,
		Description: enhancement.Description,
	}

	newE, err := cfg.Db.AddNewEnhancement(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to add new enhancement")
		return
	}

	enhancementJSON := Enhancement{
		ID:          newE.ID,
		FactionID:   newE.FactionID,
		Cost:        newE.Cost,
		Description: newE.Description,
		Detachment:  newE.Detachment,
		Name:        newE.Name,
	}

	respondWithJSON(w, 200, enhancementJSON)
}
