package controllers

import (
	"log"
	"net/http"
	"path"

	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

// NewCacheUpload : Make new cache files handle
func (state *Controller) NewCacheUpload(baseURL string) http.Handler {

	// tusd running url
	tusdURL := path.Join(baseURL, "/upload/cache")
	tusdURL += "/"

	// Tus Init
	path := state.store.GetCachePath()
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
