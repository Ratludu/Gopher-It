package main

import (
	"net/http"

	"github.com/Ratludu/Gopher-It/internal/auth"
	"github.com/Ratludu/Gopher-It/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find api key", err)
			return
		}

		user, err := cfg.DB.GetUser(r.Context(), apikey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Couldn't get user", err)
			return
		}

		handler(w, r, user)
	}
}
