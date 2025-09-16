package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"warhammer/internal/db"
)

type wargearResponse struct {
	DatasheetID int32  `json:"datasheet_id"`
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Range       string `json:"range"`
	Type        string `json:"type"`
	A           string `json:"attacks"`
	BsWs        string `json:"BS_WS"`
	Strength    string `json:"strength"`
	Ap          *int32 `json:"AP"`
	Damage      string `json:"damage"`
}

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

	wargears, err := cfg.Db.GetWargearForModel(r.Context(), int32Value)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "wargear not found")
		return
	}

	wargearSlice := []wargearResponse{}
	for _, wargear := range wargears {

		var ap *int32
		if wargear.Ap.Valid {
			ap = &wargear.Ap.Int32
		}
		wargearJSON := wargearResponse{
			DatasheetID: wargear.DatasheetID,
			Id:          wargear.ID,
			Name:        wargear.Name,
			Range:       wargear.Range,
			Type:        wargear.Type,
			A:           wargear.A,
			BsWs:        wargear.BsWs,
			Strength:    wargear.Strength,
			Ap:          ap,
			Damage:      wargear.Damage,
		}
		wargearSlice = append(wargearSlice, wargearJSON)
	}
	respondWithJSON(w, 200, wargearSlice)
}

func (cfg *ApiConfig) UpdateWargear(w http.ResponseWriter, r *http.Request) {

	type wargearRequest struct {
		DatasheetID int32  `json:"datasheet_id"`
		Id          int32  `json:"id"`
		Name        string `json:"name"`
		Range       string `json:"range"`
		Type        string `json:"type"`
		A           string `json:"attacks"`
		BsWs        string `json:"BS_WS"`
		Strength    string `json:"strength"`
		Ap          int32  `json:"AP"`
		Damage      string `json:"damage"`
	}

	request := wargearRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	fmt.Println(request)

	wargear := Wargear{
		DatasheetID: request.DatasheetID,
		Id:          request.Id,
		Name:        request.Name,
		Range:       request.Range,
		Type:        request.Type,
		A:           request.A,
		BsWs:        request.BsWs,
		Strength:    request.Strength,
		Ap:          sql.NullInt32{Int32: int32(request.Ap), Valid: true},
		Damage:      request.Damage,
	}

	paramWargear := db.UpdateWargearParams{
		DatasheetID: wargear.DatasheetID,
		ID:          wargear.Id,
		Name:        wargear.Name,
		Range:       wargear.Range,
		Type:        wargear.Type,
		A:           wargear.A,
		BsWs:        wargear.BsWs,
		Strength:    wargear.Strength,
		Ap:          wargear.Ap,
		Damage:      wargear.Damage,
	}

	updatedWargear, err := cfg.Db.UpdateWargear(r.Context(), paramWargear)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to update wargear")
		return
	}

	var ap *int32
	if updatedWargear.Ap.Valid {
		ap = &updatedWargear.Ap.Int32
	}

	wargearJSON := wargearResponse{
		DatasheetID: updatedWargear.DatasheetID,
		Id:          updatedWargear.ID,
		Name:        updatedWargear.Name,
		Range:       updatedWargear.Range,
		Type:        updatedWargear.Type,
		A:           updatedWargear.A,
		BsWs:        updatedWargear.BsWs,
		Strength:    updatedWargear.Strength,
		Ap:          ap,
		Damage:      updatedWargear.Damage,
	}

	respondWithJSON(w, 200, wargearJSON)
}

func (cfg *ApiConfig) GetWargearForModelsAll(w http.ResponseWriter, r *http.Request) {

	wargears, err := cfg.Db.GetWargearForAll(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to fetch wargear")
		return
	}
	
	wargearSlice := []Wargear{}

	for _, w := range wargears {
		wargearJSON := Wargear{
			DatasheetID: w.DatasheetID,
			Id: w.ID,
			Name: w.Name,
			Range: w.Range,
			Type: w.Type,
			A: w.A,
			BsWs: w.BsWs,
			Strength: w.Strength,
			Ap: w.Ap,
			Damage: w.Damage,
		}
		wargearSlice = append(wargearSlice, wargearJSON)
	}

	respondWithJSON(w, 200, wargearSlice)
}
