package main

import (
	"encoding/json"
	"fmt"
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

	err = cfg.db.DeleteArmy(r.Context(), id)
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

	army := ArmyList{}
	if err := json.Unmarshal(data.ArmyList, &army); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid army_list format")
		return
	}

	dbData := db.SaveToRosterParams{
		ID:           uuid.New(),
		UserID:       data.UserID,
		ArmyList:     data.ArmyList,
		Enhancements: data.Enhancement,
		Name:         data.Name,
		Faction:      data.Faction,
	}
	fmt.Println(dbData)

	err = cfg.db.SaveToRoster(r.Context(), dbData)
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

	rows, err := cfg.db.GetArmies(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch armies")
		return
	}

	resp := make([]Roster, 0, len(rows))
	for _, a := range rows {

		army := ArmyList{}
		if len(a.ArmyList) > 0 {
			if err := json.Unmarshal(a.ArmyList, &army); err != nil {
				respondWithError(w, http.StatusInternalServerError, "Malformed army_list in DB")
				return
			}
		}

		normalizeArmy(&army)

		buf, err := json.Marshal(army)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to encode army_list")
			return
		}

		resp = append(resp, Roster{
			Id:          a.ID,
			UserID:      a.UserID,
			ArmyList:    json.RawMessage(buf),
			Enhancement: a.Enhancements,
			Name:        a.Name,
			Faction:     a.Faction,
		})
	}

	respondWithJSON(w, http.StatusOK, resp)
}

func normalizeArmy(a *ArmyList) {
	if a.Character == nil {
		a.Character = []int{}
	}
	if a.Battleline == nil {
		a.Battleline = []int{}
	}
	if a.Transport == nil {
		a.Transport = []int{}
	}
	if a.Mounted == nil {
		a.Mounted = []int{}
	}
	if a.Aircraft == nil {
		a.Aircraft = []int{}
	}
	if a.Infantry == nil {
		a.Infantry = []int{}
	}
	if a.Monster == nil {
		a.Monster = []int{}
	}
	if a.Vehicle == nil {
		a.Vehicle = []int{}
	}
}
