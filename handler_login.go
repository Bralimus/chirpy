package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Bralimus/chirpy/internal/auth"
	"github.com/Bralimus/chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	dbuser, err := cfg.db.GetUser(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, dbuser.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	userToken, err := auth.MakeJWT(dbuser.ID, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unable to create token", err)
		return
	}

	refreshTokenString, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create refresh token", err)
		return
	}

	refreshToken, err := cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshTokenString,
		UserID:    dbuser.ID,
		ExpiresAt: time.Now().UTC().AddDate(0, 0, 60),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create refresh token", err)
		return
	}

	userResponse := response{
		User: User{
			ID:          dbuser.ID,
			CreatedAt:   dbuser.CreatedAt,
			UpdatedAt:   dbuser.UpdatedAt,
			Email:       dbuser.Email,
			IsChirpyRed: dbuser.IsChirpyRed},
		Token:        userToken,
		RefreshToken: refreshToken.Token,
	}

	respondWithJSON(w, http.StatusOK, userResponse)
}
