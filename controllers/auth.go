package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type formAuth struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// AuthManager : Use to Make Token JWT From Password auth
func (state *Controller) AuthManager(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		renderError(w, http.StatusBadRequest)
		return
	}

	var form formAuth

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		renderError(w, http.StatusBadRequest)
		return
	}

	if !state.upload.HasUpload(form.ID) {
		renderError(w, http.StatusNotFound)
		return
	}

	upload := state.upload.FindUpload(form.ID)

	if !upload.Auth {
		renderError(w, http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(upload.Password), []byte(form.Password)); err != nil {
		renderError(w, http.StatusUnauthorized)
		return
	}

	token, err := state.auth.GenerateNewTokenWithID(form.ID)
	if err != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}
