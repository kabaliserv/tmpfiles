package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/kabaliserv/tmpfiles/models"

	"github.com/gorilla/mux"
)

type metaRender struct {
	ID    string        `json:"id"`
	Files []models.File `json:"files"`
}

func (state *metaRender) toJSON() []byte {
	dataByte, _ := json.Marshal(state)
	return dataByte
}

// MetadataManager : Handle for get metadata upload or files
func (state *Controller) MetadataManager(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uploadID := vars["id"]

	if !state.upload.HasUpload(uploadID) {
		renderError(w, http.StatusNotFound)
		return
	}

	upload := state.upload.FindUpload(uploadID)

	if upload.IsExpire() {
		renderError(w, http.StatusNotFound)
		return
	}

	if upload.Auth {
		authorization := r.Header.Get("Authorization")
		log.Println(authorization, "toto")
		if authorization == "" {
			renderError(w, http.StatusUnauthorized)
			return
		}

		authArray := strings.Split(authorization, " ")
		if len(authArray) != 2 {
			renderError(w, http.StatusUnauthorized)
			return
		}
		if authArray[1] == "" {
			renderError(w, http.StatusUnauthorized)
			return
		}

		id := state.auth.GetIDFromToken(authArray[1])
		if id != uploadID {
			renderError(w, http.StatusUnauthorized)
			return
		}
	}

	// Make Data Render
	var dataRander = metaRender{}
	dataRander.ID = uploadID
	dataRander.Files = state.file.FindAllFileByUploadID(uploadID)

	// Send Data To JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(dataRander.toJSON())
}
