package storage

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// FileMoveFromCache : Move File in Cache To File Directory
func (state *Store) FileMoveFromCache(oldID, newID string) error {
	oldPath := filepath.Join(state.cache, oldID)
	newPath := filepath.Join(state.files, newID)
	if err := os.Rename(oldPath, newPath); err != nil {
		log.Panic(err)
	}
	return nil
}

// FileRead : Open File
func (state *Store) FileRead(id string) (*os.File, error) {
	return os.Open(filepath.Join(state.files, id))

}

// FileReadFromBytes get a byte array of the file in the file Store
func (state *Store) FileReadFromBytes(id string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(state.files, id))

}

// FileGetStat get stat file in File Store
func (state *Store) FileGetStat(id string) (os.FileInfo, error) {
	return os.Stat(filepath.Join(state.files, id))
}

// FileHasExist check if the file exists in File Store
func (state *Store) FileHasExist(id string) bool {
	if _, err := os.Stat(filepath.Join(state.files, id)); err != nil {
		return false
	}
	return true
}
