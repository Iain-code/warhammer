package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"warhammer/internal/db"

	"github.com/go-chi/chi/v5"
)

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
