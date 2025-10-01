package handlers

import (
	"net/http"

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
			Legend:      item.Legend,
			Description: item.Description,
			Field8:      item.Field8,
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
			Legend:      item.Legend,
			Description: item.Description,
			Field8:      item.Field8,
		}
		enhanceSlice = append(enhanceSlice, enhanceJSON)
	}

	respondWithJSON(w, 200, enhanceSlice)
}
