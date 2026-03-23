package main

import (
	"net/http"

	"github.com/Old-Goggles/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to retrieve ChirpID", err)
		return
	}

	rToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't retrieve token", err)
		return
	}

	userID, err := auth.ValidateJWT(rToken, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token", err)
		return
	}

	chirp, err := cfg.database.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't retrieve chirp", err)
		return
	}

	if userID != chirp.UserID {
		respondWithError(w, http.StatusForbidden, "You can't delete this chirp", nil)
		return
	}

	err = cfg.database.DeleteChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete chirp", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
