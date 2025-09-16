package handlers

import (
	"encoding/json"
	"net/http"
	"warhammer/internal/db"

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

	err = cfg.Db.DeleteArmy(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to delete army")
	}

	respondWithJSON(w, 200, "army successfully deleted")
}

func (cfg *ApiConfig) SaveToRoster(w http.ResponseWriter, r *http.Request) {

	data := Roster{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	armyJSON, err := json.Marshal(data.ArmyList)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to marshall data")
		return
	}

	dbData := db.SaveToRosterParams{
		ID:           uuid.New(),
		UserID:       data.UserID,
		ArmyList:     armyJSON,
		Enhancements: data.Enhancement,
		Name:         data.Name,
		Faction:      data.Faction,
	}

	err = cfg.Db.SaveToRoster(r.Context(), dbData)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to save roster")
		return
	}

	respondWithJSON(w, 200, "army saved successfully")
}

func (cfg *ApiConfig) GetArmies(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("user_id")
	id, err := uuid.Parse(str)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	rows, err := cfg.Db.GetArmies(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch armies")
		return
	}

	resp := []Roster{}

	for _, a := range rows {
		army := ArmyList{}
		if len(a.ArmyList) > 0 {
			if err := json.Unmarshal(a.ArmyList, &army); err != nil {
				respondWithError(w, http.StatusInternalServerError, "Malformed army_list in DB")
				return
			}
		}

		resp = append(resp, Roster{
			Id:          a.ID,
			UserID:      a.UserID,
			ArmyList:    army,
			Enhancement: a.Enhancements,
			Name:        a.Name,
			Faction:     a.Faction,
		})
	}

	respondWithJSON(w, http.StatusOK, resp)
}
