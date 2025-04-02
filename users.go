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

	// get json email / password --
	// hash password
	// add user to database
	// marshall user struct and respond with JSON
	type NewUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	newUser := NewUser{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if newUser.Password == "" || newUser.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Email and password are required")
		return
	}
	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "password invalid")
	}

	userParams := db.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
		Email:          newUser.Email,
		HashedPassword: sql.NullString{String: hashedPassword, Valid: true},
	}
	user, err := cfg.db.CreateUser(r.Context(), userParams)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to create user")
	}
	userJSON := User{
		Id:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}

	respondWithJSON(w, 200, userJSON)
}
