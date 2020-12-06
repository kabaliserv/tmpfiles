package scheduler

import (
	"github.com/kabaliserv/tmpfiles/models"
	"github.com/kabaliserv/tmpfiles/modules/clean"
	"github.com/kabaliserv/tmpfiles/storage"

	"github.com/robfig/cron/v3"
)

// Taskers struct
type Taskers struct {
	files   *models.FileManager
	uploads *models.UploadManager
	stores  *storage.Stores
}

var tasks *Taskers

// Init all scheduler
func Init() {
	tasks = &Taskers{
		files:   models.GetFileManager(),
		uploads: models.GetUploadManager(),
		stores:  storage.GetStore(),
	}

	c := cron.New()
	// Init scheduler for delete file after expire
	c.AddFunc("@every 1h", clean.CleanupUpload)

	// Start All Scheduler
	c.Start()
}
