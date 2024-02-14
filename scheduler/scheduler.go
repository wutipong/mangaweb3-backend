package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/ent"
)

var scheduler *gocron.Scheduler

type Options struct {
	EntClient *ent.Client
}

func Init(options Options) {
	scheduler = gocron.NewScheduler(time.UTC)
	scheduler.Every(30).Minutes().Do(func() {
		log.Info().Msg("Update metadata set.")
		ScanLibrary(options.EntClient)
		log.Info().Msg("Update tag list.")
		UpdateTags(options.EntClient)
		log.Info().Msg("Update missing thumbnails.")
		UpdateMissingThumbnail(options.EntClient)
	})
}

func Start() {
	scheduler.StartAsync()
}

func Stop() {
	scheduler.Stop()
}
