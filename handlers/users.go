package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if newUser.Password == "" || newUser.Username == "" {
		respondWithError(w, http.StatusBadRequest, "username and password required")
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
		HashedPassword: hashedPassword,
	}

	user, err := cfg.Db.CreateUser(r.Context(), userParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	userJSON := User{
		Id:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		Username:       user.Username,
		HashedPassword: hashedPassword,
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

	gotUser, err := cfg.Db.GetUserFromUsername(r.Context(), user.Username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}
	err = cfg.Db.DeleteUser(r.Context(), gotUser.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func (cfg *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {

	type TokenUser struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Username  string    `json:"username"`
		IsAdmin   bool      `json:"is_admin"`
		Token     string    `json:"token"`
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

	user, err := cfg.Db.GetUserFromUsername(r.Context(), newUser.Username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = auth.CompareHashedPassword(newUser.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "incorrect password")
		return
	}

	jwtToken, err := auth.MakeJWT(user.ID, cfg.TokenSecret, 15*time.Minute)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to make token")
		return
	}

	err = cfg.Db.DeleteUsersTokens(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to delete refresh tokens")
		return
	}

	tknR, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to make refresh token")
		return
	}

	refreshParams := db.CreateRefreshTokenParams{
		Token:     tknR,
		UserID:    user.ID,
		ExpiresAt: sql.NullTime{Time: time.Now().Add(24 * time.Hour), Valid: true},
	}
	_, err = cfg.Db.CreateRefreshToken(r.Context(), refreshParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to make refresh token")
		return
	}

	tknUser := TokenUser{
		ID:        user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  user.Username,
		IsAdmin:   user.IsAdmin,
		Token:     jwtToken,
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    tknR,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,                // change to true after dev
		SameSite: http.SameSiteLaxMode, // change Lax => Strict after dev
		Path:     "/",
	})
	respondWithJSON(w, 200, tknUser)
}

func (cfg *ApiConfig) RefreshHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "missing refresh token")
		return
	}

	tkn, err := cfg.Db.GetRefreshToken(r.Context(), cookie.Value)
	if err != nil || !tkn.ExpiresAt.Valid || tkn.ExpiresAt.Time.Before(time.Now()) {
		respondWithError(w, http.StatusUnauthorized, "refresh token invalid or expired")
		return
	}

	err = cfg.Db.DeleteRefreshToken(r.Context(), tkn.Token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to delete token")
		return
	}

	newRefreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to make refresh token")
		return
	}

	refreshParams := db.CreateRefreshTokenParams{
		Token:     newRefreshToken,
		UserID:    tkn.UserID,
		ExpiresAt: sql.NullTime{Time: time.Now().Add(24 * time.Hour), Valid: true},
	}

	rTkn, err := cfg.Db.CreateRefreshToken(r.Context(), refreshParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to add refresh token to database")
		return
	}

	jwtToken, err := auth.MakeJWT(tkn.UserID, cfg.TokenSecret, 15*time.Minute)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to make JWT token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    rTkn.Token,
		Expires:  rTkn.ExpiresAt.Time,
		HttpOnly: true,
		Secure:   false,                // change to true after dev
		SameSite: http.SameSiteLaxMode, // change Lax => Strict after dev
		Path:     "/",
	})

	respondWithJSON(w, 200, map[string]string{
		"token": jwtToken,
	})
}

func (cfg *ApiConfig) MakeAdmin(w http.ResponseWriter, r *http.Request) {

	user := User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}
	fmt.Println("this is user data")
	fmt.Println(user)

	dbUser, err := cfg.Db.GetUser(r.Context(), user.Id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "invalid username")
		return
	}

	_, err = cfg.Db.MakeAdmin(r.Context(), dbUser.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to make admin")
		return
	}

	respondWithJSON(w, 200, "admin rights granted")
}

func (cfg *ApiConfig) RemoveAdmin(w http.ResponseWriter, r *http.Request) {

	user := User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	dbUser, err := cfg.Db.GetUser(r.Context(), user.Id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "invalid username")
		return
	}

	_, err = cfg.Db.RemoveAdmin(r.Context(), dbUser.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to remove admin tag")
		return
	}

	respondWithJSON(w, 200, "admin rights removed")
}
