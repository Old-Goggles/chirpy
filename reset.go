package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Forbidden", nil)
		return
	}

	cfg.fileserverHits.Store(0)
	err := cfg.database.DeleteUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not reset users", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
