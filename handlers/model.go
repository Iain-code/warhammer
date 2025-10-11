package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"warhammer/internal/db"
)

func (cfg *ApiConfig) GetModel(w http.ResponseWriter, r *http.Request) {

	datasheetID := r.URL.Query().Get("datasheet_id")
	if datasheetID == "" {
		respondWithError(w, http.StatusBadRequest, "Datasheet ID not provided")
		return
	}

	datasheetIDInt64, err := strconv.ParseInt(datasheetID, 0, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid datasheet_id")
		return
	}
	datasheetIDInt32 := int32(datasheetIDInt64)

	model, err := cfg.Db.GetModel(r.Context(), datasheetIDInt32)
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

func (cfg *ApiConfig) GetAllModels(w http.ResponseWriter, r *http.Request) {

	models, err := cfg.Db.GetAllModels(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to get models")
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

func (cfg *ApiConfig) GetModelsForFaction(w http.ResponseWriter, r *http.Request) {

	factionID := r.URL.Query().Get("faction_id")
	if factionID == "" {
		respondWithError(w, http.StatusBadRequest, "Datasheet ID not provided")
		return
	}
	str := sql.NullString{String: factionID, Valid: true}

	models, err := cfg.Db.GetModelsForFaction(r.Context(), str)
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

func (cfg *ApiConfig) UpdateModel(w http.ResponseWriter, r *http.Request) {

	model := Model{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&model)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	paramModel := db.UpdateModelParams{
		DatasheetID: model.DatasheetID,
		OldID:       model.OldID,
		Name:        model.Name,
		M:           model.M,
		T:           model.T,
		W:           model.W,
		Sv:          model.Sv,
		InvSv:       model.InvSv,
		Ld:          model.Ld,
		Oc:          model.Oc,
	}

	updatedModel, err := cfg.Db.UpdateModel(r.Context(), paramModel)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to update model")
		return
	}

	modelJSON := Model{
		OldID:       updatedModel.OldID,
		DatasheetID: updatedModel.DatasheetID,
		Name:        updatedModel.Name,
		M:           updatedModel.M,
		T:           updatedModel.T,
		Sv:          updatedModel.Sv,
		InvSv:       updatedModel.InvSv,
		W:           updatedModel.W,
		Ld:          updatedModel.Ld,
		Oc:          updatedModel.Oc,
	}

	respondWithJSON(w, 200, modelJSON)
}

func (cfg *ApiConfig) AddNewModel(w http.ResponseWriter, r *http.Request) {

	model := Model{}

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&model)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	
}
