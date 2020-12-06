package models

import (
	"encoding/json"
	"log"
	"sync"

	nanoid "github.com/aidarkhanov/nanoid/v2"
)

// File struct
type File struct {
	ID           uint   `gorm:"primaryKey" json:"-"`
	UploadID     string `gorm:"not null" json:"-"`
	FileID       string `gorm:"uniqueIndex;not null" json:"id"`
	Name         string `gorm:"not null" json:"name"`
	Type         string `json:"type"`
	Size         int    `gorm:"not null" json:"size"`
	LastModified int64  `json:"-"`
}

// FileManager struct
type FileManager struct {
	db *DB
}

var filemgr *FileManager

// GetFileManager : Make File Manager if is not create and return this
func GetFileManager() *FileManager {
	var once sync.Once
	once.Do(func() {
		GetDB().AutoMigrate(&File{})
		filemgr = &FileManager{}
		filemgr.db = database
	})
	return filemgr
}

// HasFile : Check if the given fileID exists.
func (state *FileManager) HasFile(fileID string) bool {
	row, err := state.db.Where("file_id=?", fileID).Find(&File{}).Rows()
	if err != nil {
		return false
	}
	defer row.Close()
	if row.Next() {
		return true

	}
	return false
}

// FindFile : Find File By ID
func (state *FileManager) FindFile(fileID string) *File {
	file := File{}
	state.db.Where("file_id=?", fileID).Find(&file)
	return &file
}

// FindAllFileByUploadID : Find All Files By UploadID
func (state *FileManager) FindAllFileByUploadID(uploadID string) []File {
	files := []File{}

	rows, err := state.db.Model(&File{}).Where("upload_id = ?", uploadID).Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var file File
		// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
		state.db.ScanRows(rows, &file)

		files = append(files, file)

		// do something
	}

	return files
}

// AddFile :
func (state *FileManager) AddFile(file *File) error {
	for {
		fileID, err := nanoid.GenerateString(nanoid.DefaultAlphabet, 12)
		if err != nil {
			log.Println(err)
		}
		if !state.HasFile(fileID) {
			file.FileID = fileID
			break
		}
	}

	state.db.Create(file)

	return nil
}

// ToJSON seralize File Struct to json
func (state *File) ToJSON() []byte {
	dataByte, _ := json.Marshal(state)
	return dataByte
}
