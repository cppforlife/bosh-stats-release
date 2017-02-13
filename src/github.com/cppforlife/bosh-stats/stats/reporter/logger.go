package reporter

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type Logger struct {
	logTag string
	logger boshlog.Logger
}

func NewLogger(logger boshlog.Logger) Logger {
	return Logger{"stats.reporter.Logger", logger}
}

func (r Logger) Report(stats []Stat) error {
	for _, stat := range stats {
		r.logger.Debug(r.logTag, `name="%s" value="%s" tags="%#v"`, stat.Name(), stat.Value(), stat.Tags())
	}

	return nil
}
