package main

import (
	"net/http"

	"github.com/jcsmurph/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevokeToken(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT")
		return
	}

	revokeErr := cfg.DB.RevokeToken(token)

	if revokeErr != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to revoke token")
		return
	}

	respondWithStatus(w, http.StatusOK)
}
