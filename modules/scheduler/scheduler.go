package scheduler

import (
	"github.com/kabaliserv/tmfiles/models"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	DB *models.DB
}

// Start is function for exe all scheduler
func Start() {
	c := cron.New()
	// Init scheduler for delete file after expire
	c.AddFunc("@every 1h", removeAllExpiredFiles)

	// Start All Scheduler
	c.Start()
}
