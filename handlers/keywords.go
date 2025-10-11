package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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
