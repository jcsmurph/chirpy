package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jcsmurph/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerChirpDelete(w http.ResponseWriter, r *http.Request) {


	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT")
		return
	}

	validAccessToken := auth.ValidateAccessToken(token, cfg.jwtSecret)
	if validAccessToken != nil {
		respondWithError(w, http.StatusUnauthorized, "Token is not an access token")
		return
	}

	subject, ValidateJWTerr := auth.ValidateJWT(token, cfg.jwtSecret)
	if ValidateJWTerr != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT")
		return
	}

    chirpIDString := chi.URLParam(r, "chirpID")
    chirpID, err := strconv.Atoi(chirpIDString)
    userID, err := strconv.Atoi(subject)

    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Chirp ID")
        return
    }

    dbChirp, err := cfg.DB.GetChirpID(chirpID)

    if err != nil {
        respondWithError(w, http.StatusNotFound, "Couldn't get chirp")
        return
    }

    deleteChirpErr := cfg.DB.DeleteChirp(dbChirp.ID, userID)

    if deleteChirpErr != nil {
        respondWithError(w, http.StatusForbidden, "The user is not authorized to delete the chirp")
    }

	respondWithJSON(w, http.StatusOK, Chirp{})

}
