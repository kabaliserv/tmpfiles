package models

import (
	"time"

	nanoid "github.com/aidarkhanov/nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

// Upload struct
type Upload struct {
	ID        uint   `gorm:"primaryKey"`
	UploadID  string `gorm:"uniqueIndex" json:"id"`
	Auth      bool
	Password  string
	CreatedAt time.Time
	ExpireAt  time.Time
	Permanent bool
}

// UploadManager struct
type UploadManager struct {
	db *DB
}

// NewUploadManager : Create new *UploadManager that can be used for managing uploads data.
func NewUploadManager(db *DB) (*UploadManager, error) {
	db.AutoMigrate(&Upload{})

	uploadmgr := UploadManager{}

	uploadmgr.db = db

	return &uploadmgr, nil
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

// FindAllUpload : Find all Uploads
func (state *UploadManager) FindAllUpload() *[]Upload {
	uploads := []Upload{}
	state.db.Model(&Upload{}).Find(&uploads)

	// Sanitize Uplads with remove Password
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
			panic(err)
		}
		if state.HasUpload(uploadID) {
			continue
		}

		upload.UploadID = uploadID
		break
	}

	if upload.Password != "" {
		upload.Auth = true
		pwd := upload.Password
		upload.Password = state.HashPassword(pwd)
	}
	state.db.Create(&upload)
	return nil
}

// HashPassword : Hash the password (takes a upload as well, it can be used for salting).
func (state *UploadManager) HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unsuccessful")
	}
	return string(hash)
}

// CheckPassword : compare a hashed password with a possible plaintext equivalent
func (state *UploadManager) CheckPassword(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
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
