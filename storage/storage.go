package storage

import (
	"log"
	"os"
	"path/filepath"
)

// Store struct
type Store struct {
	data, files, cache string
}

// NewStore : Make Store with path directory
func NewStore(path string) (*Store, error) {

	// Get absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Print(err)
	}

	if _, err := os.Stat(absPath); err != nil {
		return nil, err
		// log.Fatal("Directory: ", absPath, " NOT FOUND")
	}

	var gerr error

	data := filepath.Join(absPath, "data")
	if _, err := os.Stat(data); err != nil {
		gerr = os.Mkdir(data, 0700)
	}

	files := filepath.Join(absPath, "files")
	if _, err := os.Stat(files); err != nil {
		gerr = os.Mkdir(files, 0700)
	}

	cache := filepath.Join(absPath, "cache")
	if _, err := os.Stat(cache); err != nil {
		gerr = os.Mkdir(cache, 0700)
	}

	if gerr != nil {
		return nil, gerr
	}

	return &Store{
		data:  data,
		files: files,
		cache: cache,
	}, nil
}

// GetDataPath :
func (state *Store) GetDataPath() string {
	return state.data
}

// GetFilesPath :
func (state *Store) GetFilesPath() string {
	return state.files
}

// GetCachePath :
func (state *Store) GetCachePath() string {
	return state.cache
}
