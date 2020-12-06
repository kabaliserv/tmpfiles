package controllers

import (
	"log"
	"net/http"
	"path"

	"github.com/kabaliserv/tmpfiles/config"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

// InitTusServer : return new tus handler
func (state *Controller) InitTusServer() http.Handler {

	// tusd running url
	tusdURL := path.Join(config.GetWebPath(), "/api/upload/cache")
	tusdURL += "/"

	// Tus Init
	path := state.stores.GetCachePath()
	store := filestore.New(path)

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	config := tusd.Config{
		BasePath:      tusdURL,
		StoreComposer: composer,
	}

	handler, err := tusd.NewHandler(config)
	if err != nil {
		log.Fatal(err)
	}

	return http.StripPrefix(tusdURL, handler)
}
