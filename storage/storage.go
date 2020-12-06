package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kabaliserv/tmpfiles/config"

	"golang.org/x/sys/unix"
)

// Stores struct
type Stores struct {
	data, files, cache string
}

var path Stores

// Init : Make Store with config path directory
func Init() error {

	if !dirIsWritable(config.GetStorePath()) {
		return fmt.Errorf("Store directory is not writable")
	}

	path = Stores{
		data:  ifIsNotExistMakeDirectory("data"),
		files: ifIsNotExistMakeDirectory("files"),
		cache: ifIsNotExistMakeDirectory("cache"),
	}
	errText := "directory is not writable"
	if !dirIsWritable(path.data) {
		return fmt.Errorf("data ", errText)
	}
	if !dirIsWritable(path.files) {
		return fmt.Errorf("files ", errText)
	}
	if !dirIsWritable(path.cache) {
		return fmt.Errorf("cache ", errText)
	}

	return nil
}

func ifIsNotExistMakeDirectory(path string) string {
	dir := filepath.Join(config.GetStorePath(), path)
	if _, err := os.Stat(dir); err != nil {
		_ = os.Mkdir(dir, 0700)
	}
	return dir
}

func dirIsWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

// GetStore :
func GetStore() *Stores {
	return &path
}

// GetDataPath :
func (state *Stores) GetDataPath() string {
	return state.data
}

// GetFilesPath :
func (state *Stores) GetFilesPath() string {
	return state.files
}

// GetCachePath :
func (state *Stores) GetCachePath() string {
	return state.cache
}

// deleteFileFromStore : deletion according to store
func (state *Stores) deleteFileFromStore(store, fileName string) error {
	var path string
	switch store {
	case "cache":
		path = filepath.Join(state.cache, fileName)
	case "files":
		path = filepath.Join(state.files, fileName)
	}
	log.Println(path)
	if err := os.Remove(path); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
