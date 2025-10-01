package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"warhammer/internal/db"
	"strings"

	"github.com/go-chi/chi/v5"
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

func (cfg *ApiConfig) GetKeywordsForFaction(w http.ResponseWriter, r *http.Request) {

	factionID := r.URL.Query().Get("faction_id")
	if factionID == "" {
		respondWithError(w, http.StatusBadRequest, "Keyword ID not provided")
		return
	}
	str := sql.NullString{String: factionID, Valid: true}

	models, err := cfg.Db.GetModelsForFaction(r.Context(), str)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "models not found")
		return
	}

	modelSlice := []int32{}
	for _, model := range models {
		modelSlice = append(modelSlice, model.DatasheetID)
	}

	keywords, err := cfg.Db.GetKeywordsForFaction(r.Context(), modelSlice)
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

	idsStr := chi.URLParam(r, "ids")
	idsStr = strings.TrimSpace(idsStr)
    if idsStr == "" {
        respondWithError(w, http.StatusBadRequest, "points_id is required")
        return
    }

	parts := strings.Split(idsStr, ",")

	parsed := make([]int32, 0, len(parts))
	for _, s := range parts {
		n, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "invalid points_id: "+s)
			return
		}
		parsed = append(parsed, int32(n))
	}

	points, err := cfg.Db.GetPointsForID(r.Context(), parsed)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to get points")
		return
	}

	out := make([]Points, 0, len(points))
	for _, p := range points {
		out = append(out, Points{
			Id:          p.ID,
			DatasheetID: p.DatasheetID,
			Line:        p.Line,
			Description: p.Description,
			Cost:        p.Cost,
		})
	}
	respondWithJSON(w, http.StatusOK, out)
}

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

func (cfg *ApiConfig) UpdatePoints(w http.ResponseWriter, r *http.Request) {

	Id := chi.URLParam(r, "id")
	if Id == "" {
		respondWithError(w, http.StatusBadRequest, "ID not provided")
		return
	}

	pointsModel := Points{}
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&pointsModel)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	parseId, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to parse Id")
		return
	}

	Id32 := int32(parseId)

	model, err := cfg.Db.GetPointsForOneID(r.Context(), Id32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to find id model")
		return
	}

	params := db.UpdatePointsForIDParams{
		ID:          model.ID,
		DatasheetID: model.DatasheetID,
		Line:        model.Line,
		Description: model.Description,
		Cost:        pointsModel.Cost,
	}

	fmt.Printf("params: %v\n", params)
	fmt.Printf("Id: %v\n", Id)
	fmt.Printf("Model: %v\n", model)

	updatedPoints, err := cfg.Db.UpdatePointsForID(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to update points")
		return
	}

	pointsJSON := Points{
		Id:          updatedPoints.ID,
		DatasheetID: updatedPoints.DatasheetID,
		Line:        updatedPoints.Line,
		Description: updatedPoints.Description,
		Cost:        updatedPoints.Cost,
	}

	respondWithJSON(w, 200, pointsJSON)
}

func (cfg *ApiConfig) GetKeywordsForModel(w http.ResponseWriter, r *http.Request) {
	Id := chi.URLParam(r, "id")

	parsedID, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to parse Id")
		return
	}

	Id32 := int32(parsedID)

	models, err := cfg.Db.GetKeywordsForModel(r.Context(), Id32)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to fetch keywords")
		return
	}

	keywords := KeywordsModel{
		Id:          models[1].ID,
		DatasheetID: models[1].DatasheetID,
	}

	for _, model := range models {
		if model.Keyword != "" {
			keywords.Keywords = append(keywords.Keywords, model.Keyword)
		}
	}

	respondWithJSON(w, 200, keywords)

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
