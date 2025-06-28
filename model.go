package main

import (
	"database/sql"
	"net/http"
)

func (cfg *ApiConfig) GetModel(w http.ResponseWriter, r *http.Request) {

	datasheetID := r.URL.Query().Get("datasheet_id")
	if datasheetID == "" {
		respondWithError(w, http.StatusBadRequest, "Datasheet ID not provided") // Ensure this isn't the response
		return
	}

	model, err := cfg.db.GetModel(r.Context(), datasheetID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "model not found")
		return
	}

	modelJSON := Model{
		OldID:       model.OldID,
		DatasheetID: model.DatasheetID,
		Name:        model.Name,
		M:           model.M,
		T:           model.T,
		Sv:          model.Sv,
		InvSv:       model.InvSv,
		W:           model.W,
		Ld:          model.Ld,
		Oc:          model.Oc,
	}
	respondWithJSON(w, 200, modelJSON)

}

func (cfg *ApiConfig) GetModelsForFaction(w http.ResponseWriter, r *http.Request) {

	factionID := r.URL.Query().Get("faction_id")
	if factionID == "" {
		respondWithError(w, http.StatusBadRequest, "Faction ID not provided")
		return
	}
	str := sql.NullString{String: factionID, Valid: true}

	models, err := cfg.db.GetModelsForFaction(r.Context(), str)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "model not found")
		return
	}
	modelSlice := []Model{}
	for _, model := range models {
		modelJSON := Model{
			OldID:       model.OldID,
			DatasheetID: model.DatasheetID,
			Name:        model.Name,
			M:           model.M,
			T:           model.T,
			Sv:          model.Sv,
			InvSv:       model.InvSv,
			W:           model.W,
			Ld:          model.Ld,
			Oc:          model.Oc,
		}
		modelSlice = append(modelSlice, modelJSON)
	}
	respondWithJSON(w, 200, modelSlice)
}
