package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Old-Goggles/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerWebhook(w http.ResponseWriter, r *http.Request) {
	type polkaData struct {
		User_ID string `json:"user_id"`
	}

	type polkaParams struct {
		Event string    `json:"event"`
		Data  polkaData `json:"data"`
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing Api Key", err)
		return
	}
	if apiKey != cfg.polkakey {
		respondWithError(w, http.StatusUnauthorized, "Invalid API Key", nil)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := polkaParams{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	userID, err := uuid.Parse(params.Data.User_ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to parse user id", err)
		return
	}

	_, err = cfg.database.SetChirpyRed(r.Context(), userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "User not found", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Couldn't upgrade user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
