package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/kabaliserv/tmpfiles/models"
)

// FormUpload struct
type formUpload struct {
	Auth     bool     `json:"auth"`
	Password string   `json:"password"`
	Expire   int      `json:"expire"` // (0 = 5min), (1 = 10min), (2 = 1h), (3 = 6h), (4 = 1j), (5 = 3j)
	FilesID  []string `json:"filesid"`
}

// UploadManager : Add Upload Handle
func (state *Controller) UploadManager(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("content-type") != "application/json" {
		renderError(w, http.StatusBadRequest)
		return
	}

	var form formUpload

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		renderError(w, http.StatusBadRequest)
		return
	}

	for _, v := range form.FilesID {
		if !state.stores.CacheFileHasExist(v) {
			renderError(w, http.StatusBadRequest)
			return
		}
	}

	upload := &models.Upload{}
	if form.Auth {
		upload.Auth = true
		upload.Password = form.Password
	}
	upload.ExpireAt = makeTimeWithExpireData(form.Expire)

	if err := state.upload.AddUpload(upload); err != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}

	for _, v := range form.FilesID {
		file, err := state.stores.CacheGetMeta(v)
		if err != nil {
			renderError(w, http.StatusInternalServerError)
			return
		}

		file.UploadID = upload.UploadID
		if err := state.file.AddFile(file); err != nil {
			renderError(w, http.StatusInternalServerError)
			return
		}

		if err := state.stores.FileMoveFromCache(v, file.FileID); err != nil {
			renderError(w, http.StatusInternalServerError)
			return
		}

		if err := state.stores.CacheRemoveFileInfo(v); err != nil {
			renderError(w, http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte(upload.UploadID))
}

func makeTimeWithExpireData(expire int) time.Time {

	switch expire {
	case 1:
		return time.Now().Add(time.Duration(10) * time.Minute)
	case 2:
		return time.Now().Add(time.Duration(1) * time.Hour)
	case 3:
		return time.Now().Add(time.Duration(6) * time.Hour)
	case 4:
		return time.Now().Add(time.Duration(1) * time.Hour * 24)
	case 5:
		return time.Now().Add(time.Duration(3) * time.Hour * 24)
	}

	return time.Now().Add(time.Duration(5) * time.Minute)
}
