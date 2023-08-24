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
		log.Info().Msg("Update metadata set.")
		ScanLibrary()
		log.Info().Msg("Update tag list.")
		UpdateTags()
		log.Info().Msg("Update missing thumbnails.")
		UpdateMissingThumbnail()
	})
}

func Start() {
	scheduler.StartAsync()
}

func Stop() {
	scheduler.Stop()
}
