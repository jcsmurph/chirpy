package main

import (
	"net/http"
	"strconv"

	"github.com/jcsmurph/chirpy/internal/auth"
)

type AccessToken struct {
	AccessToken string `json:"token"`
}

func (cfg *apiConfig) handlerRefreshToken(w http.ResponseWriter, r *http.Request) {
	type response struct {
		AccessToken
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT")
		return
	}

	revokeErr := cfg.DB.CheckRevokeToken(token)
	if revokeErr != nil {
		respondWithError(w, http.StatusUnauthorized, "Token has been revoked")
	}

	subject, validateErr := auth.ValidateJWT(token, cfg.jwtSecret)
	if validateErr != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT")
		return
	}

    idInt, err := strconv.Atoi(subject)

    if err != nil {
        respondWithError(w, http.StatusUnauthorized, "Unable to convert ID from string to integer")
		return
    }

    accessToken, _, err := auth.MakeJWT(idInt, cfg.jwtSecret)

	respondWithJSON(w, http.StatusOK, response{
		AccessToken: AccessToken{
			AccessToken: accessToken,
		},
	})
}
