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
	Password string     `json:"password"`
	Expire   ExpireData `json:"expire"`
	FilesID  []string   `json:"filesid"`
}

// ExpireData struct
type ExpireData struct {
	Time string `json:"time"`
	Val  int    `json:"val"`
}

// PostUploads : Add Upload Handle
func (state *Controller) PostUploads(w http.ResponseWriter, r *http.Request) {
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

	upload := &models.Upload{}

	upload.Password = form.Password
	upload.ExpireAt = makeTimeWithExpireData(&form.Expire)

	if err := state.upload.AddUpload(upload); err != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}

	for _, v := range form.FilesID {
		file, err := state.store.CacheGetMeta(v)
		if err != nil {
			renderError(w, http.StatusInternalServerError)
			return
		}

		file.UploadID = upload.UploadID
		if err := state.file.AddFile(file); err != nil {
			renderError(w, http.StatusInternalServerError)
			return
		}

		if err := state.store.FileMoveFromCache(v, file.FileID); err != nil {
			renderError(w, http.StatusInternalServerError)
			return
		}

		if err := state.store.CacheRemoveFileInfo(v); err != nil {
			renderError(w, http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte(upload.UploadID))
}

func makeTimeWithExpireData(expire *ExpireData) time.Time {
	type minMax struct {
		Min int
		Max int
	}
	var conf = map[string]minMax{
		"minute": minMax{Min: 5, Max: 59},
		"hour":   minMax{Min: 1, Max: 23},
		"day":    minMax{Min: 1, Max: 7},
	}

	if conf[expire.Time] == (minMax{0, 0}) || expire.Val < conf[expire.Time].Min || expire.Val > conf[expire.Time].Max {
		return time.Now().Add(5 * time.Minute)
	}

	switch expire.Time {
	case "minute":
		return time.Now().Add(time.Duration(expire.Val) * time.Minute)
	case "hour":
		return time.Now().Add(time.Duration(expire.Val) * time.Hour)
	case "day":
		return time.Now().Add(time.Duration(expire.Val) * 24 * time.Hour)
	}

	return time.Now().Add(5 * time.Minute)
}
