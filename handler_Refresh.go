package main

import (
	"net/http"
	"time"

	"github.com/Old-Goggles/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	rToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't retrieve token", err)
		return
	}

	user, err := cfg.database.GetUserFromRefreshToken(r.Context(), rToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couln't retrieve user", err)
		return
	}

	newToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't make JWT", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: newToken,
	})
}
