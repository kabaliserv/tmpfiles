package clean

import (
	"log"

	"github.com/kabaliserv/tmpfiles/models"
	"github.com/kabaliserv/tmpfiles/storage"
)

// CleanupUpload : 
func CleanupUpload() {
	var countUpload int
	var countFiles int
	uploads := models.GetUploadManager().FindAllActiveUploads()
	chanCountFilesRemove := make(chan int)
	for _, v := range *uploads {
		if !v.IsExpire() {
			continue
		}
		files := models.GetFileManager().FindAllFileByUploadID(v.UploadID)
		go removeAllExpiredFiles(chanCountFilesRemove, files)
		models.GetUploadManager().DisableUpload(v.UploadID)
		countUpload++
		filesRemove := <-chanCountFilesRemove
		countFiles += filesRemove
	}
	log.Println(countUpload, "Upload Disable and", countFiles, "Files Remove")
}


func removeAllExpiredFiles(chanCount chan int, files []models.File) {
	count := 0
	for _, v := range files {
		if err := storage.GetStore().FileDelete(v.FileID); err == nil {
			count++
		}
	}
	chanCount <- count
}