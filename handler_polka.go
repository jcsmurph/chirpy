package main

import (
	"encoding/json"
	"net/http"

	"github.com/jcsmurph/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUpgradeUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
        Data map[string]int `json:"data"`
        Event string `json:"event"`
    }


   polkaKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Polka API key is missing")
		return
	}

    if polkaKey != cfg.polkaSecret {
        respondWithError(w, http.StatusUnauthorized, "Polka key is incorrect")
    }

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	decodeErr := decoder.Decode(&params)
	if decodeErr != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters for User Update handler")
		return
	}

	user, err := cfg.DB.GetUserID(params.Data["user_id"])

	if err != nil {
        respondWithError(w, http.StatusNotFound, "Unable to find User")
		return
	}

	if params.Event != "user.upgraded" {
		respondWithStatus(w, http.StatusOK)
		return
	}

	upgradeErr := cfg.DB.UpgradeUser(user.ID)

	if upgradeErr != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to upgrade user")
		return
	}

    respondWithStatus(w, http.StatusOK)

}
