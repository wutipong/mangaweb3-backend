package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

var scheduler *gocron.Scheduler

type Options struct {
}

func Init(options Options) {
	scheduler = gocron.NewScheduler(time.UTC)
	scheduler.Every(30).Minutes().Do(func() {
		log.Print("Update metadata set.")
		ScanLibrary()
		log.Print("Update tag list.")
		UpdateTags()
		log.Print("Update missing thumbnails.")
		UpdateMissingThumbnail()
	})
}

func Start() {
	scheduler.StartAsync()
}

func Stop() {
	scheduler.Stop()
}
