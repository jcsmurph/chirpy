package main

import (
	"encoding/json"
	"net/http"

	"github.com/jcsmurph/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
	}
	type response struct {
		User
		AccessToken string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
    decoder := json.NewDecoder(r.Body)
	params := parameters{}
    err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters for User Update handler")
		return
	}

	user, err := cfg.DB.GetUserByEmail(params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get user")
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	accessToken, refreshToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create JWT")
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:    user.ID,
			Email: user.Email,
            Redchirpy: user.RedChirpy,
		},
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	})
}
