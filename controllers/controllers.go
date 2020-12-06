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
	stores  *storage.Stores
	auth   *auth.Data
}

var manager *Controller

// Init : create new controller for use handle with http
func Init() {
	authmgr := &auth.Data{Secret: []byte("hzbeeursygfrhbxiugqyeibfqfdhsqkjuhfd")}
	manager =  &Controller{
		upload: models.GetUploadManager(),
		file:   models.GetFileManager(),
		stores:  storage.GetStore(),
		auth:   authmgr,
	}
}

// GetManagers :
func GetManagers() *Controller {
	return manager
}