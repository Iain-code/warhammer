package main

import (
	"net/http"
	"time"
	"warhammer/internal/auth"
)

func (cfg *ApiConfig) Refresh(w http.ResponseWriter, r *http.Request) {

	type Token struct {
		Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find header")
		return
	}
	user, err := cfg.db.GetUserFromToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "invalid user")
		return
	}
	jwtToken, err := auth.MakeJWT(user.ID, cfg.tokenSecret, 1*time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to make JWT token")
		return
	}
	_, err = auth.ValidateJWT(jwtToken, cfg.tokenSecret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to validate JWT token")
		return
	}
	tkn := Token{
		Token: jwtToken,
	}
	respondWithJSON(w, http.StatusOK, tkn)
}

func (cfg *ApiConfig) Revoke(w http.ResponseWriter, r *http.Request) {

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid header")
		return
	}
	err = cfg.db.RevokeRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to remove token")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
