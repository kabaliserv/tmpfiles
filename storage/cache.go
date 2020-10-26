package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kabaliserv/tmpfiles/models"
)

type filedata struct {
	Size     int       `json:"Size"`
	Metadata *metadata `json:"MetaData"`
}

type metadata struct {
	Name         string `json:"filename"`
	Type         string `json:"filetype"`
	LastModified string `json:"lastmodified"`
}

// CacheGetMeta : Get meta from Cache directory
func (state *Store) CacheGetMeta(id string) (*models.File, error) {
	data := filedata{}
	path := id + ".info"
	src, err := state.cacheGetReaderFile(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("popo")
	err = json.Unmarshal(src, &data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	file := &models.File{}
	file.Name = data.Metadata.Name
	file.Type = data.Metadata.Type
	file.Size = data.Size
	lastmodified, err := strconv.Atoi(data.Metadata.LastModified)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	file.LastModified = int64(lastmodified)

	return file, nil
}

// Get Reader From file in Cache Directory
func (state *Store) cacheGetReaderFile(id string) ([]byte, error) {
	fileReader, err := ioutil.ReadFile(filepath.Join(state.cache, id))
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return fileReader, nil
}

// CacheFileHasExist : Get if file exist in Cache Store
func (state *Store) CacheFileHasExist(id string) bool {
	if _, err := os.Stat(filepath.Join(state.cache, id)); err != nil {
		return false
	}
	return true
}

// CacheRemoveFileInfo : Remove file info to cache directory
func (state *Store) CacheRemoveFileInfo(id string) error {
	file := id + ".info"
	path := filepath.Join(state.cache, file)
	if err := deleteFileToCache(path); err != nil {
		return err
	}
	return nil
}

func deleteFileToCache(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
