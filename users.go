package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
	"warhammer/internal/auth"
	"warhammer/internal/db"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	type NewUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	newUser := NewUser{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if newUser.Password == "" || newUser.Username == "" {
		respondWithError(w, http.StatusBadRequest, "Email and password are required")
		return
	}
	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "password invalid")
		return
	}

	userParams := db.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
		Username:       newUser.Username,
		HashedPassword: sql.NullString{String: hashedPassword, Valid: true},
	}
	user, err := cfg.db.CreateUser(r.Context(), userParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	userJSON := User{
		Id:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
	}

	respondWithJSON(w, 200, userJSON)
}

func (cfg *ApiConfig) DeleteUser(w http.ResponseWriter, r *http.Request) {

	user := User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	gotUser, err := cfg.db.GetUserFromEmail(r.Context(), user.Username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}
	err = cfg.db.DeleteUser(r.Context(), gotUser.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func (cfg *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {

	// decode JSON
	// check users hashed password
	// find user by email
	// check JWT token
	// get refresh token

	type TokenUser struct {
		ID           uuid.UUID `json:"id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Username     string    `json:"username"`
		IsAdmin      bool      `json:"is_admin"`
		Token        string    `json:"token"`
		RefreshToken string    `json:"refresh_token"`
	}
	type NewUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	newUser := NewUser{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}
	user, err := cfg.db.GetUserFromEmail(r.Context(), newUser.Username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}
	err = auth.CompareHashedPassword(newUser.Password, user.HashedPassword.String)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "incorrect password")
		return
	}

	jwtToken, err := auth.MakeJWT(user.ID, cfg.tokenSecret, 15*time.Minute)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to make token")
		return
	}

	userID, err := auth.ValidateJWT(jwtToken, cfg.tokenSecret)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid token")
		return
	}

	tknR, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to make refresh token")
		return
	}

	refreshParams := db.CreateRefreshTokenParams{
		Token:     tknR,
		UserID:    userID,
		ExpiresAt: sql.NullTime{Time: time.Now().Add(24 * time.Hour), Valid: true},
	}
	_, err = cfg.db.CreateRefreshToken(r.Context(), refreshParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to make refresh token")
		return
	}
	tknUser := TokenUser{
		ID:           userID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Username:     user.Username,
		IsAdmin:      user.IsAdmin,
		Token:        jwtToken,
		RefreshToken: tknR,
	}

	respondWithJSON(w, 200, tknUser)
}
