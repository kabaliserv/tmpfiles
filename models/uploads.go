package models

import (
	"fmt"
	"time"
	"sync"

	nanoid "github.com/aidarkhanov/nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

// Upload struct
type Upload struct {
	ID        uint   `gorm:"primaryKey"`
	UploadID  string `gorm:"uniqueIndex" json:"id"`
	Active    bool   `gorm:"default:true"`
	Auth      bool
	Password  string
	CreatedAt time.Time
	ExpireAt  time.Time
	Permanent bool
}

// UploadManager struct
type UploadManager struct {
	db DB
}

var uploadmgr *UploadManager

// GetUploadManager : Make Upload Manager if is not create and return this
func GetUploadManager() *UploadManager {
	var once sync.Once
	once.Do(func() {
		GetDB().AutoMigrate(&Upload{})
		uploadmgr = &UploadManager{}
		uploadmgr.db = *database
	})
	return uploadmgr
}

// HasUpload : Check if the given uploadID exists.
func (state *UploadManager) HasUpload(uploadID string) bool {
	if err := state.db.Where("upload_id=?", uploadID).First(&Upload{}).Error; err != nil {
		return false
	}
	return true
}

// FindUpload : Find Upload By ID
func (state *UploadManager) FindUpload(uploadID string) *Upload {
	upload := Upload{}
	state.db.Where("upload_id=?", uploadID).Find(&upload)
	return &upload
}

// FindAllUploads : Find all Uploads
func (state *UploadManager) FindAllUploads() *[]Upload {
	uploads := []Upload{}
	state.db.Model(&Upload{}).Find(&uploads)

	// Sanitize Uploads
	for _, v := range uploads {
		v.Password = ""
	}

	return &uploads
}

// FindAllActiveUploads : get all upload active
func (state *UploadManager) FindAllActiveUploads() *[]Upload {
	uploads := []Upload{}

	state.db.Model(&Upload{}).Where("active=?", true).Find(&uploads)

	// Sanitize Uploads
	for _, v := range uploads {
		v.Password = ""
	}

	return &uploads
}

// AddUpload :
func (state *UploadManager) AddUpload(upload *Upload) error {
	for {
		uploadID, err := nanoid.GenerateString(nanoid.DefaultAlphabet, 9)
		if err != nil {
			return err
		}
		if state.HasUpload(uploadID) {
			continue
		}
		upload.UploadID = uploadID
		break
	}
	if upload.Auth {
		if err := upload.HashPassword(); err != nil {
			return err
		}
	}
	state.db.Create(&upload)
	return nil
}

// DisableUpload : 
func (state *UploadManager) DisableUpload(id string) error {
	state.db.Model(&Upload{}).Where("upload_id=?", id).Update("active", false)
	return nil
}

// HashPassword : Hash the password (takes a upload as well, it can be used for salting).
func (state *Upload) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(state.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%v %v", "Permissions: bcrypt password hashing unsuccessful ", err)
	}
	state.Password = string(hash)
	return nil
}

// CheckPassword : compare a hashed password with a possible plaintext equivalent
func (state *Upload) CheckPassword(password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(state.Password), []byte(password)) != nil {
		return false
	}
	return true
}

// IsExpire check if Upload has expire
func (state *Upload) IsExpire() bool {
	if time.Now().After(state.ExpireAt) {
		return true
	}
	return false
}

//
