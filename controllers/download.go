package controllers

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
)

// DownloadManager :
func (state *Controller) DownloadManager(w http.ResponseWriter, r *http.Request) {
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
		keys, ok := r.URL.Query()["t"]
		log.Printf("%#v", keys)
		if !ok || len(keys[0]) < 1 {
			renderError(w, http.StatusUnauthorized)
			return
		}
		token := keys[0]

		id := state.auth.GetIDFromToken(token)
		if id != uploadID {
			renderError(w, http.StatusUnauthorized)
			return
		}
	}

	if keys, ok := r.URL.Query()["f"]; ok && len(keys[0]) > 0 {
		fileID := keys[0]
		if !state.file.HasFile(fileID) {
			renderError(w, http.StatusNotFound)
			return
		}

		file := state.file.FindFile(fileID)

		if !state.stores.FileHasExist(fileID) {
			renderError(w, http.StatusNotFound)
			return
		}
		src, err := state.stores.FileRead(fileID)
		if err != nil {
			renderError(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", file.Type)
		w.Header().Set("Content-Length", fmt.Sprintf("%v", file.Size))
		contentDisposition := "attachment; filename*=UTF-8''" + url.QueryEscape(file.Name)
		w.Header().Set("Content-Disposition", contentDisposition)
		w.WriteHeader(http.StatusOK)
		io.Copy(w, src)
		return
	}

	files := state.file.FindAllFileByUploadID(uploadID)

	for _, file := range files {
		if !state.stores.FileHasExist(file.FileID) {
			renderError(w, http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("content-type", "application/zip")
	contentDisposition := "attachment; filename*=UTF-8''TMPFiles-" + time.Now().Format("02-01-2006") + ".zip"
	w.Header().Set("Content-Disposition", contentDisposition)
	w.WriteHeader(http.StatusOK)

	buffZip := zip.NewWriter(w)

	for _, file := range files {
		src, err := state.stores.FileReadFromBytes(file.FileID)
		if err != nil {
			log.Println(err)
		}
		stat, err := state.stores.FileGetStat(file.FileID)
		if err != nil {
			log.Println(err)
		}

		zipHeader, err := zip.FileInfoHeader(stat)
		if err != nil {
			log.Println(err)
		}
		zipHeader.Name = file.Name
		zipHeader.Modified = time.Unix(0, (file.LastModified * int64(time.Millisecond)))
		zipHeader.Flags = 0x800
		///h := &zip.FileHeader{Name: file.FileName, Method: zip.Deflate, Modified: time.Now(), Flags: 0x800}
		ziper, err := buffZip.CreateHeader(zipHeader)
		if err != nil {
			log.Println(err)
		}

		_, err = ziper.Write(src)
		if err != nil {
			log.Println(err)
		}
	}

	if err := buffZip.Close(); err != nil {
		log.Println(err)
	}
}
