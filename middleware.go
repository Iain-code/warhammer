package main

import (
	"net/http"
	"warhammer/internal/auth"
)

func (cfg *ApiConfig) middlewareAuth(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "no authorization header found")
			return
		}

		_, err = auth.ValidateJWT(token, cfg.tokenSecret)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "invalid token")
			return
		}
		next.ServeHTTP(w, r)
	})
}
