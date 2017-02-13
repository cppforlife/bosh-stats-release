package stats

import (
	"time"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/cppforlife/bosh-stats/stats/reporter"
)

type Worker struct {
	src      Source
	reporter reporter.Reporter

	logTag string
	logger boshlog.Logger
}

func NewWorker(src Source, reporter reporter.Reporter, logger boshlog.Logger) Worker {
	return Worker{src, reporter, "stats.Worker", logger}
}

func (w Worker) Send() error {
	ticker := time.NewTicker(1 * time.Minute) // todo add random interval?

	select {
	case <-ticker.C:
		w.send()
	}

	return nil
}

func (w Worker) send() {
	stats, err := w.src.Stats()
	if err != nil {
		w.logger.Error(w.logTag, "Failed to get stats from '%T': %s", w.src, err)
	}

	if len(stats) > 0 {
		var repStats []reporter.Stat

		for _, stat := range stats {
			repStats = append(repStats, stat)
		}

		err := w.reporter.Report(repStats)
		if err != nil {
			w.logger.Error(w.logTag, "Failed to send stats: %s", err)
		}
	}
}
