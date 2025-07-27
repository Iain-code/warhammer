package main

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

	model, err := cfg.db.GetModel(r.Context(), datasheetIDInt32)
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
		respondWithError(w, http.StatusBadRequest, "Datasheet ID not provided")
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

	updatedModel, err := cfg.db.UpdateModel(r.Context(), paramModel)
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

func (cfg *ApiConfig) GetKeywordsForFaction(w http.ResponseWriter, r *http.Request) {

	factionID := r.URL.Query().Get("faction_id")
	if factionID == "" {
		respondWithError(w, http.StatusBadRequest, "Keyword ID not provided")
		return
	}
	str := sql.NullString{String: factionID, Valid: true}

	models, err := cfg.db.GetModelsForFaction(r.Context(), str)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "models not found")
		return
	}

	modelSlice := []int32{}
	for _, model := range models {
		modelSlice = append(modelSlice, model.DatasheetID)
	}

	keywords, err := cfg.db.GetKeywordsForFaction(r.Context(), modelSlice)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "keywords not found")
		return
	}

	keywordSlice := []Keyword{}

	for _, keyword := range keywords {
		keyword := Keyword{
			Id:          keyword.ID,
			DatasheetID: keyword.DatasheetID,
			Keyword:     keyword.Keyword,
		}
		keywordSlice = append(keywordSlice, keyword)
	}

	respondWithJSON(w, 200, keywordSlice)
}

func (cfg *ApiConfig) GetPointsForModels(w http.ResponseWriter, r *http.Request) {

	queryData := r.URL.Query()["points_id[]"]
	if len(queryData) == 0 {
		respondWithError(w, http.StatusBadRequest, "Keyword ID not provided")
		return
	}

	parsedArr := []int32{}
	for _, item := range queryData {
		parsed, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "failed to parse data")
		}
		parsedInt32 := int32(parsed)
		parsedArr = append(parsedArr, parsedInt32)
	}

	points, err := cfg.db.GetPointsForID(r.Context(), parsedArr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to get points")
		return
	}
	pointSlice := []Points{}

	for _, point := range points {
		JSONpoints := Points{
			Id:           point.DatasheetID,
			Datasheet_id: point.DatasheetID,
			Line:         point.Line,
			Description:  point.Description,
			Cost:         point.Cost,
		}
		pointSlice = append(pointSlice, JSONpoints)
	}

	respondWithJSON(w, 200, pointSlice)
}

func (cfg *ApiConfig) GetEnhancements(w http.ResponseWriter, r *http.Request) {

	enhance, err := cfg.db.GetEnhancements(r.Context())
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
