package controllers

import (
	"github.com/kabaliserv/tmpfiles/auth"
	"github.com/kabaliserv/tmpfiles/models"
	"github.com/kabaliserv/tmpfiles/storage"
)

// Controller struct
type Controller struct {
	upload *models.UploadManager
	file   *models.FileManager
	store  *storage.Store
	auth   *auth.Data
}

// NewController : Get new controller for use handle with http
func NewController(store *storage.Store, db *models.DB) *Controller {

	uploadmgr, _ := models.NewUploadManager(db)
	filemgr, _ := models.NewFileManager(db)

	authmgr := &auth.Data{Secret: []byte("hzbeeursygfrhbxiugqyeibfqfdhsqkjuhfd")}

	return &Controller{
		upload: uploadmgr,
		file:   filemgr,
		store:  store,
		auth:   authmgr,
	}
}
