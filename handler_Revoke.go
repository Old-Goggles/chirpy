package main

import (
	"net/http"

	"github.com/Old-Goggles/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	rToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't retrieve token", err)
		return
	}

	_, err = cfg.database.RevokeRefreshToken(r.Context(), rToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
