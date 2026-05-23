package worker

import (
	"log"

	"aac-hls-stream-prep/internal/track"
)

type Worker struct {
	id      int
	service *track.Service
	jobs    chan Job
}

func NewWorker(
	id int,
	service *track.Service,
	jobs chan Job,
) *Worker {
	return &Worker{
		id:      id,
		service: service,
		jobs:    jobs,
	}
}

func (w *Worker) Start() {
	go func() {
		for job := range w.jobs {

			log.Printf(
				"[worker %d] processing track %s",
				w.id,
				job.TrackID,
			)

			err := w.service.ProcessTrack(
				job.TrackID,
				job.OriginalPath,
			)

			if err != nil {
				log.Printf(
					"[worker %d] failed processing %s: %v",
					w.id,
					job.TrackID,
					err,
				)
				continue
			}

			log.Printf(
				"[worker %d] finished track %s",
				w.id,
				job.TrackID,
			)
		}
	}()
}
