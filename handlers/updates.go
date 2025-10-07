package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"warhammer/internal/db"

	"github.com/go-chi/chi/v5"
)

func (cfg *ApiConfig) UpdateAbility(w http.ResponseWriter, r *http.Request) {
	Id := chi.URLParam(r, "id")
	Line := chi.URLParam(r, "line")

	parsedId, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to parse Id")
		return
	}

	parsedLine, err := strconv.ParseInt(Line, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to parse Id")
		return
	}

	Id32 := int32(parsedId)
	Line32 := int32(parsedLine)

	AbilityUpdate := AbilityUpdate{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&AbilityUpdate)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	abilityParams := db.GetAbilityParams{
		DatasheetID: Id32,
		Line:        Line32,
	}

	ability, err := cfg.Db.GetAbility(r.Context(), abilityParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to get ability")
		return
	}

	params := db.UpdateAbilitiesParams{
		DatasheetID: Id32,
		Line:        Line32,
		AbilityID:   ability.AbilityID,
		Model:       ability.Model,
		Name:        ability.Name,
		Description: AbilityUpdate.Description,
		Type:        ability.Type,
		Parameter:   ability.Parameter,
	}

	err = cfg.Db.UpdateAbilities(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to update ability")
		return
	}

	s := []string{}
	s = append(s, ability.Name)
	s = append(s, AbilityUpdate.Description)

	respondWithJSON(w, 200, s)
}

func (cfg *ApiConfig) DeleteUnit(w http.ResponseWriter, r *http.Request) {
	Id := chi.URLParam(r, "id")

	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to parse Id")
		return
	}

	Id32 := int32(Id64)

	err = cfg.Db.DeleteUnitFromModels(r.Context(), Id32)
	if err != nil {
		fmt.Printf("delete error: %v", err)
		respondWithError(w, http.StatusBadRequest, "unable to delete unit from models")
		return
	}

	respondWithJSON(w, 200, "Unit removed successfully")
}

func (cfg *ApiConfig) UpdateEnhancements(w http.ResponseWriter, r *http.Request) {

	enhanceUpdate := EnhancementUpdate{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&enhanceUpdate)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	Id := chi.URLParam(r, "id")

	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to parse Id")
		return
	}

	Id32 := int32(Id64)

	enhanceDb, err := cfg.Db.GetEnhancementFromId(r.Context(), Id32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to find enhancement")
		return
	}

	cost := enhanceDb.Cost
	description := enhanceDb.Description

	if enhanceUpdate.Cost != nil {
		cost = *enhanceUpdate.Cost
	}

	if enhanceUpdate.Description != nil {
		description = *enhanceUpdate.Description
	}

	params := db.UpdateEnhancementParams{
		ID:          Id32,
		Cost:        int32(cost),
		Description: description,
		Detachment:  enhanceUpdate.Detachment,
		FactionID:   enhanceUpdate.FactionID,
		Name:        enhanceUpdate.Name,
	}

	err = cfg.Db.UpdateEnhancement(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to update enhancement")
		return
	}

	respondWithJSON(w, 200, "enhancement updated successfully")
}

func (cfg *ApiConfig) DeleteEnhancements(w http.ResponseWriter, r *http.Request) {
	Id := chi.URLParam(r, "id")

	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to parse Id")
		return
	}

	Id32 := int32(Id64)

	err = cfg.Db.DeleteEnhancement(r.Context(), Id32)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to delete enhancement")
		return
	}

	respondWithJSON(w, 200, "enhancement successfully deleted")
}
