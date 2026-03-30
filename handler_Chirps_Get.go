package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("author_id")
	if authorID == "" {
		dbChirps, err := cfg.database.GetAllChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
			return
		}

		chirps := []Chirp{}
		for _, chirp := range dbChirps {
			chirps = append(chirps, Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				UserID:    chirp.UserID,
				Body:      chirp.Body,
			})
		}

		respondWithJSON(w, http.StatusOK, chirps)
	} else {
		parsedAuthorID, err := uuid.Parse(authorID)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Unable to parse author ID", err)
			return
		}

		chirpsByAuthor, err := cfg.database.GetChirpsByAuthorID(r.Context(), parsedAuthorID)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Couldn't retrieve chirps", err)
			return
		}

		chirps := []Chirp{}
		for _, chirp := range chirpsByAuthor {
			chirps = append(chirps, Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				UserID:    chirp.UserID,
				Body:      chirp.Body,
			})
		}

		respondWithJSON(w, http.StatusOK, chirps)
	}
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("chirpID")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to parse ID", err)
		return
	}
	dbChirp, err := cfg.database.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't retrieve chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		UserID:    dbChirp.UserID,
		Body:      dbChirp.Body,
	})
}
