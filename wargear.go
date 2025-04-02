package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (cfg *ApiConfig) GetWargearForModel(w http.ResponseWriter, r *http.Request) {
	datasheetID := r.URL.Query().Get("datasheet_id")
	if datasheetID == "" {
		respondWithError(w, http.StatusBadRequest, "Datasheet ID not provided")
		return
	}

	value, err := strconv.ParseInt(datasheetID, 10, 32)
	if err != nil {
		fmt.Println("Error while parsing:", err)
		return
	}

	int32Value := int32(value)

	wargears, err := cfg.db.GetWargearForModel(r.Context(), int32Value)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "wargear not found")
		return
	}
	wargearSlice := []Wargear{}
	for _, wargear := range wargears {
		wargearJSON := Wargear{
			DatasheetID: wargear.DatasheetID,
			Field2:      wargear.Field2,
			Name:        wargear.Name,
			Range:       wargear.Range,
			Type:        wargear.Type,
			A:           wargear.A,
			BsWs:        wargear.BsWs,
			Strength:    wargear.Strength,
			Ap:          wargear.Ap,
			Damage:      wargear.Damage,
		}
		wargearSlice = append(wargearSlice, wargearJSON)
	}
	respondWithJSON(w, 200, wargearSlice)
}

// why do we use r.context() for the context
// why is the input a string for wargear model
